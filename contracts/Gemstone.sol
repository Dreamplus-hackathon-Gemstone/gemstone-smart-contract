// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC1155/extensions/ERC1155Burnable.sol";
import "@openzeppelin/contracts/token/ERC1155/extensions/ERC1155Supply.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

contract GEM is ERC1155, Ownable, ERC1155Burnable, ERC1155Supply {
    using Counters for Counters.Counter;

    Counters.Counter private _tokenIdCounter;

    /* Events */
    event MintProposal(uint256 tokenId, address indexed maker);

    event MintMovie(uint256 tokenId, address indexed maker);

    event FundingSuccess(
        uint256 tokenId,
        address indexed maker,
        Proposal proposal
    );

    uint256 public constant MAKER_REVENUE_RATIO = 50;
    uint256 public constant MINER_REVENUE_RATIO = 40;
    uint256 public constant GEMSTONE_REVENUE_RATIO = 10;

    uint256 private FeeRevenues;

    constructor() ERC1155("") {
        _tokenIdCounter.increment();
    }

    struct Miner {
        uint256 amount;
        address minerAddress;
    }

    struct Proposal {
        uint256 tokenId;
        uint256 targetAmount;
        uint256 Deadline;
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

    function _setURI(uint256 tokenId, string memory tokenURI) internal {
        URIStorage[tokenId] = tokenURI;
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

    function invest(uint256 tokenId)
        public
        payable
        isApproved
        isLocked(tokenId)
    {
        // 투자 기간 확인
        // 잔량/송금액 확인
        // mint
    }

    function withdraw(uint256 tokenId)
        public
        payable
        isApproved
        isLocked(tokenId)
    {
        // 투자자가 맞는지 확인
        // 가지고 있는양 확인
        // maker에게 해당 양만큼 소유권 이전
        // 스토리지 amount 동기화
        // 송금
    }

    function fundingSuccess(uint256 tokenId) public isApproved {
        // 기한에 도달했는지 확인
        // msg.sender가 maker가 맞는지 확인
        // Lock 걸어놓기...
    }

    function burnProposal(uint256 tokenId) public isApproved {
        // msg.sender가 maker가 맞는지 확인
        // 남아있는 금액 확인
        // 전체 반환
        // burn NFT
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
