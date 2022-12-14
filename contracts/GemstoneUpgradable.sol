// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/token/ERC1155/ERC1155Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC1155/extensions/ERC1155BurnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
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

contract GemstoneUpgradable is
    Initializable,
    ERC1155Upgradeable,
    OwnableUpgradeable,
    ERC1155BurnableUpgradeable,
    UUPSUpgradeable
{
    using Counters for Counters.Counter;

    Counters.Counter private _tokenIdCounter;

    /* Events */
    event MintProposal(uint256 tokenId, address indexed maker);

    event MintMovie(uint256 tokenId, address indexed maker);

    event MintAgenda(uint256 tokenId, address indexed miner);

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

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
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

    struct Voter {
        uint256 weight;
        uint256 vote;
        bool voted;
    }

    // tokenId와 proposal 정보 매핑
    mapping(uint256 => Proposal) ProposalMapping;

    // proposal의 잠금 상태 확인
    mapping(uint256 => bool) ProposalLockStatus;

    // tokenId => 주소 => Miner ... 개인 투자자의 정보를 확인하거나, 투자자가 맞는 지 확인할 때 사용
    mapping(uint256 => mapping(address => Miner)) Investors;

    // token URI Storage
    mapping(uint256 => string) URIStorage;

    modifier isLocked(uint256 tokenId) {
        require(ProposalLockStatus[tokenId] == false);

        ProposalLockStatus[tokenId] = true;

        _;

        ProposalLockStatus[tokenId] = false;
    }

    modifier isApproved() {
        require(
            isApprovedForAll(msg.sender, address(this)) == true,
            "Caller have to approve this contract address managing asset"
        );

        _;
    }

    function initialize() public initializer {
        __Ownable_init();
        __UUPSUpgradeable_init();
        USDc = USDC(0x0FA8781a83E46826621b3BC094Ea2A0212e71B23);
    }

    function _authorizeUpgrade(address newImplementation)
        internal
        override
        onlyOwner
    {}

    /* tokenId에 해당하는 Proposal의 정보를 확인합니다. */
    function viewProposal(uint256 tokenId)
        public
        view
        isApproved
        returns (Proposal memory)
    {
        return ProposalMapping[tokenId];
    }

    /* msg.sender의 tokenId에 대한 투자 정보를 확인합니다. */
    function viewInvestorInfo(uint256 tokenId)
        public
        view
        isApproved
        returns (Miner memory)
    {
        return Investors[tokenId][msg.sender];
    }

    /* 투자자 배열을 확인합니다. */
    function viewInvestorList(uint256 tokenId)
        public
        view
        isApproved
        returns (Miner[] memory)
    {
        return ProposalMapping[tokenId].investorList;
    }

    /* USDC의 잔액을 확인합니다. */
    function balanceOfUSDC() public view returns (uint256) {
        return USDc.balanceOf(msg.sender);
    }

    /*
     * User의 Gemstone 컨트랙트에 대한 USDC Allowance를 확인합니다. 해당 값 만큼 User는 컨트랙트에 송금할 수 있습니다.
     * 해당 값이 모자라는 경우에는 직접 USDC 컨트랙트를 호출하여 approve함수로 값을 증가시켜야 합니다.
     */
    function checkAllowance() public view returns (uint256) {
        return USDc.allowance(msg.sender, address(this));
    }

    /*
     * 유저가 회원가입 할 때, 해당 함수를 호출해서 컨트랙트가 유저에게 USDC를 보낼 수 있는 한도를 uint256의 최댓값으로 설정합니다.
     * 유저가 요청할 때 마다 스토리지를 확인하거나 값을 변경하는 작업은 추가적인 가스를 요하기 때문에, 미리 최댓값으로 설정하도록 했습니다.
     */
    function approveUSDC() public {
        USDc.approve(
            msg.sender,
            115792089237316195423570985008687907853269984665640564039457584007913129639935
        );
    }

    /* 해당 tokenId의 tokenURI를 관리하는 함수입니다. URIStorage라는 mapping 자료구조에 값을 넣어 저장합니다. */
    function _setURI(uint256 tokenId, string memory tokenURI) internal {
        URIStorage[tokenId] = tokenURI;
    }

    /* funding에 성공했을 경우, 혹은 특정 상황(아직 미정) 해당 proposal에 lock을 걸어줍니다. */
    function lockProposal(uint256 tokenId) internal isApproved {
        ProposalLockStatus[tokenId] = true;
    }

    /* NFT minting, dest 주소로 amount개를 minting하고, tokenURI를 세팅합니다. */
    function mintNFT(
        address dest,
        uint256 amount,
        string memory tokenURI
    ) internal returns (uint256) {
        // tokenID 증가
        _tokenIdCounter.increment();
        uint256 newTokenId = _tokenIdCounter.current();

        // mint
        _mint(dest, newTokenId, amount, "");
        // setURI
        _setURI(newTokenId, tokenURI);

        return newTokenId;
    }

    /* Proposal을 mint합니다. */
    function mintProposal(uint256 amount, string memory tokenURI)
        internal
        returns (uint256)
    {
        uint256 tokenId = mintNFT(address(this), amount, tokenURI);

        emit MintProposal(tokenId, address(this));

        return tokenId;
    }

    function sendUSDC(
        address _to,
        address _from,
        uint256 $USDC
    ) public payable isApproved {
        // 잔량/송금액 확인
        require(balanceOfUSDC() >= $USDC, "Caller don't have enough USDC");

        // Allowance 확인
        require(
            USDc.allowance(_from, _to) >= $USDC * 10**6,
            "Allowance is lower than caller claimed. Please check allowance."
        );

        USDc.transferFrom(_from, _to, $USDC * 10**6);
    }

    /* Proposal을 업로드 합니다. */
    function propose(
        uint256 targetAmount,
        uint256 deadline,
        string calldata tokenURI
    ) public payable isApproved {
        // 최소 target amount
        require(
            targetAmount > 10000000000,
            "Minimum target amount is 10,000$, Check your target amount"
        );

        // tokenID 증가
        _tokenIdCounter.increment();

        // tokenID
        uint256 newTokenId = _tokenIdCounter.current();

        // Proposal mint
        mintProposal(targetAmount, tokenURI);

        // Storage 갱신
        ProposalMapping[newTokenId].tokenId = newTokenId;
        ProposalMapping[newTokenId].targetAmount = targetAmount;
        ProposalMapping[newTokenId].deadline = deadline;
        ProposalMapping[newTokenId].currentFunded = 0;
        ProposalMapping[newTokenId].makerAddress = msg.sender;

        ProposalLockStatus[newTokenId] = false;
    }

    /* 투자 시 호출하는 함수 */
    function invest(uint256 tokenId, uint256 $USDC)
        public
        payable
        isApproved
        isLocked(tokenId)
    {
        Proposal memory proposal = ProposalMapping[tokenId];

        // 투자 기간 확인
        require(block.timestamp <= proposal.deadline, "The deadline is over");

        // 최대 금액 확인
        require(
            proposal.currentFunded + $USDC <= proposal.targetAmount,
            "Too much value"
        );

        // 송금
        sendUSDC(address(this), msg.sender, $USDC);

        // 소유권 이전
        safeTransferFrom(address(this), msg.sender, tokenId, $USDC, "");

        // Storage data 갱신
        Miner memory newMiner = Miner($USDC, msg.sender);
        ProposalMapping[tokenId].currentFunded += $USDC;
        ProposalMapping[tokenId].investorList.push(newMiner);
        Investors[tokenId][msg.sender] = newMiner;

        // 이벤트 발생
        emit Deposit(tokenId, msg.sender, $USDC);
    }

    /* 투자 철회시 사용하는 함수 */
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
        require(value > 0);

        // contract로 해당 양만큼 소유권 이전
        safeTransferFrom(msg.sender, address(this), tokenId, value, "");

        // 송금
        sendUSDC(msg.sender, address(this), value);

        // 스토리지 amount 동기화
        ProposalMapping[tokenId].currentFunded -= value;

        // 이벤트 발생
        emit Withdraw(tokenId, miner.minerAddress, value);
    }

    /* funding 성공 시 호출하는 함수 */
    function fundingSuccess(uint256 tokenId) public isApproved {
        Proposal memory proposal = ProposalMapping[tokenId];
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

        // target amount가 달성되었는지 확인
        require(
            proposal.currentFunded >= proposal.targetAmount,
            "Target Amount of proposal is not fulfilled."
        );

        // Lock 걸어놓기...
        lockProposal(tokenId);
    }

    /* Proposal이 시간이 다 찼지만 funding에 실패한 경우, 혹은 모종의 이유로 취소되는 경우 호출하는 함수 */
    function burnProposal(uint256 tokenId) public isApproved {
        Proposal memory proposal = ProposalMapping[tokenId];

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

    /* 투자자, 제작자 계좌에 판매 수익을 저장합니다. */
    function StoreRevenue(uint256 tokenId) internal {}

    /* 계좌에 에치된 USDC를 출금합니다. */
    function withdrawUSDC(uint256 tokenId, address sender) external onlyOwner {}
}
