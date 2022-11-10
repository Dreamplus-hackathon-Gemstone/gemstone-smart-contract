// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC1155/extensions/ERC1155Burnable.sol";
import "@openzeppelin/contracts/token/ERC1155/extensions/ERC1155Supply.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

interface USDC {
    function balanceOf(address account) external view returns (uint256);

    function allowance(address owner, address spender)
        external
        view
        returns (uint256);

    function transfer(address recipient, uint256 amount)
        external
        returns (bool);

    function approve(address spender, uint256 amount) external returns (bool);

    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) external returns (bool);
}

contract GEM is ERC1155, Ownable, ERC1155Burnable, ERC1155Supply {
    using Counters for Counters.Counter;

    Counters.Counter private _tokenIdCounter;

    /* Events */
    event MintProposal(uint256 tokenId, address indexed maker);

    event MintMovie(uint256 tokenId, address indexed maker);

    event Deposit(uint256 tokenId, address indexed miner, uint256 $USDC);

    event Withdraw(uint256 tokenId, address indexed miner, uint256 $USDC);

    event FundingSuccess(
        uint256 tokenId,
        address indexed maker,
        Proposal proposal
    );

    uint256 public constant MAKER_REVENUE_RATIO = 50;
    uint256 public constant MINER_REVENUE_RATIO = 40;
    uint256 public constant GEMSTONE_REVENUE_RATIO = 10;

    uint256 private FeeRevenues;

    USDC public USDc;

    constructor() ERC1155("") {
        USDc = USDC(0x0FA8781a83E46826621b3BC094Ea2A0212e71B23);
        _tokenIdCounter.increment();
    }

    struct Miner {
        uint256 amount;
        address minerAddress;
    }

    struct Proposal {
        uint256 tokenId;
        uint256 targetAmount;
        uint256 deadline;
        uint256 currentFunded;
        address makerAddress;
        Miner[] investorList;
    }

    mapping(uint256 => Proposal) Proposals;

    mapping(uint256 => bool) ProposalLocked;

    mapping(uint256 => mapping(address => Miner)) Investors;

    mapping(uint256 => string) URIStorage;

    modifier isLocked(uint256 tokenId) {
        require(ProposalLocked[tokenId] == false);

        ProposalLocked[tokenId] = true;

        _;

        ProposalLocked[tokenId] = false;
    }

    modifier isApproved() {
        require(
            isApprovedForAll(msg.sender, address(this)) == true,
            "Caller have to approve this contract address managing asset"
        );

        _;
    }

    function viewProposal(uint256 tokenId)
        public
        view
        isApproved
        returns (Proposal memory)
    {
        return Proposals[tokenId];
    }

    function viewInvestorInfo(uint256 tokenId)
        public
        view
        isApproved
        returns (Miner memory)
    {
        return Investors[tokenId][msg.sender];
    }

    function viewInvestorList(uint256 tokenId)
        public
        view
        isApproved
        returns (Miner[] memory)
    {
        return Proposals[tokenId].investorList;
    }

    function balanceofUSDC() public view returns (uint256) {
        return USDc.balanceOf(msg.sender);
    }

    function checkAllowance() public view returns (uint256) {
        return USDc.allowance(msg.sender, address(this));
    }

    function _setURI(uint256 tokenId, string memory tokenURI) internal {
        URIStorage[tokenId] = tokenURI;
    }

    function lockProposal(uint256 tokenId) internal isApproved {
        ProposalLocked[tokenId] = true;
    }

    function mintNFT(uint256 amount, string memory tokenURI) public isApproved {
        _tokenIdCounter.increment();
        uint256 newTokenId = _tokenIdCounter.current();
        _mint(msg.sender, newTokenId, amount, "");
        _setURI(newTokenId, tokenURI);
    }

    function mintProposal(uint256 amount, string memory tokenURI)
        public
        payable
        isApproved
    {
        mintNFT(amount, tokenURI);
    }

    function sendUSDC(
        address _to,
        address _from,
        uint256 $USDC
    ) public payable isApproved {
        // 잔량/송금액 확인
        require(balanceofUSDC() >= $USDC, "Caller don't have enough USDC");

        // Allowance 확인
        require(
            USDc.allowance(_from, _to) >= $USDC * 10**6,
            "Allowance is lower than caller claimed. Please check allowance."
        );

        USDc.transferFrom(_from, _to, $USDC * 10**6);
    }

    function invest(uint256 tokenId, uint256 $USDC)
        public
        payable
        isApproved
        isLocked(tokenId)
    {
        Proposal memory proposal = Proposals[tokenId];

        // 투자 기간 확인
        require(block.timestamp <= proposal.deadline, "The deadline is over");

        // 송금
        sendUSDC(address(this), msg.sender, $USDC);

        // 소유권 이전
        safeTransferFrom(proposal.makerAddress, msg.sender, tokenId, $USDC, "");

        // 이벤트 발생
        emit Deposit(tokenId, msg.sender, $USDC);

        // Storage data 갱신
        Proposals[tokenId].currentFunded += $USDC;
    }

    function withdraw(uint256 tokenId)
        public
        payable
        isApproved
        isLocked(tokenId)
    {
        Miner memory miner = Investors[tokenId][msg.sender];

        // 투자자가 맞는지 확인
        require(miner.minerAddress != address(0));

        // 가지고 있는양 확인
        uint256 value = miner.amount;

        // maker에게 해당 양만큼 소유권 이전
        safeTransferFrom(
            msg.sender,
            Proposals[tokenId].makerAddress,
            tokenId,
            value,
            ""
        );

        // 송금
        sendUSDC(msg.sender, address(this), value);

        // 이벤트 발생
        emit Withdraw(tokenId, miner.minerAddress, value);

        // 스토리지 amount 동기화
        Proposals[tokenId].currentFunded -= value;
    }

    function fundingSuccess(uint256 tokenId) public isApproved {
        Proposal memory proposal = Proposals[tokenId];
        // 기한에 도달했는지 확인
        require(
            block.timestamp <= proposal.deadline,
            "The deadline is not over."
        );

        // msg.sender가 maker가 맞는지 확인
        require(
            msg.sender == proposal.makerAddress,
            "Only maker can execute this function."
        );

        // Lock 걸어놓기...
        lockProposal(tokenId);
    }

    function burnProposal(uint256 tokenId) public isApproved {
        Proposal memory proposal = Proposals[tokenId];

        // msg.sender가 maker가 맞는지 확인
        require(
            msg.sender == proposal.makerAddress,
            "Only maker can execute this function."
        );

        // 남아있는 금액 확인
        uint256 value = proposal.currentFunded;

        // 전체 반환
        // burn NFT
        if (value != 0) {
            Miner[] memory miners = proposal.investorList;

            for (uint256 i = 0; i < miners.length; i++) {
                Miner memory m = miners[i];
                payable(m.minerAddress).transfer(m.amount);
                burn(m.minerAddress, tokenId, m.amount);
            }
        }
    }

    // The following functions are overrides required by Solidity.
    function _beforeTokenTransfer(
        address operator,
        address from,
        address to,
        uint256[] memory ids,
        uint256[] memory amounts,
        bytes memory data
    ) internal override(ERC1155, ERC1155Supply) {
        super._beforeTokenTransfer(operator, from, to, ids, amounts, data);
    }
}
