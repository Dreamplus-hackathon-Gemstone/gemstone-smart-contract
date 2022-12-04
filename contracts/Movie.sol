// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721BurnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/CountersUpgradeable.sol";

contract MovieContract is
    Initializable,
    ERC721Upgradeable,
    ERC721EnumerableUpgradeable,
    ERC721URIStorageUpgradeable,
    PausableUpgradeable,
    OwnableUpgradeable,
    ERC721BurnableUpgradeable,
    UUPSUpgradeable
{
    using CountersUpgradeable for CountersUpgradeable.Counter;

    CountersUpgradeable.Counter private _tokenIdCounter;

    event MintMovie(uint256 tokenId, address indexed maker);

    uint256 public constant MAKER_REVENUE_RATIO = 50;
    uint256 public constant MINER_REVENUE_RATIO = 40;
    uint256 public constant GEMSTONE_REVENUE_RATIO = 10;

    USDC public USDc;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    enum Status {
        invalid,
        unavailable,
        available
    }

    struct Movie {
        uint256 price;
        uint256 totalRevenue;
        address maker;
        string tokenURI;
    }

    // 영화
    mapping(uint256 => Movie) MovieTable;

    // 영화 구매자 목록
    mapping(uint256 => mapping(address => Status)) BuyerTable;

    // token URI Storage
    mapping(uint256 => string) URIStorage;

    function initialize() public initializer {
        __ERC721_init("Movie", "GMOVIE");
        __ERC721Enumerable_init();
        __ERC721URIStorage_init();
        __Pausable_init();
        __Ownable_init();
        __ERC721Burnable_init();
        __UUPSUpgradeable_init();
    }

    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }

    function safeMint(address to, string memory uri) public onlyOwner {
        uint256 tokenId = _tokenIdCounter.current();
        _tokenIdCounter.increment();
        _safeMint(to, tokenId);
        _setTokenURI(tokenId, uri);
    }

    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 tokenId,
        uint256 batchSize
    )
        internal
        override(ERC721Upgradeable, ERC721EnumerableUpgradeable)
        whenNotPaused
    {
        super._beforeTokenTransfer(from, to, tokenId, batchSize);
    }

    function _authorizeUpgrade(address newImplementation)
        internal
        override
        onlyOwner
    {}

    // The following functions are overrides required by Solidity.

    function _burn(uint256 tokenId)
        internal
        override(ERC721Upgradeable, ERC721URIStorageUpgradeable)
    {
        super._burn(tokenId);
    }

    function tokenURI(uint256 tokenId)
        public
        view
        override(ERC721Upgradeable, ERC721URIStorageUpgradeable)
        returns (string memory)
    {
        return super.tokenURI(tokenId);
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721Upgradeable, ERC721EnumerableUpgradeable)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }

    /* USDC 컨트랙트를 정의합니다. */
    function setUSDC(address contractAddr) external onlyOwner {
        USDc = USDC(contractAddr);
    }

    /* 영화 정보를 확인합니다. */
    function getMovie(uint256 tokenId) public view returns (Movie memory) {
        return MovieTable[tokenId];
    }

    /* 구매 여부를 확인합니다. */
    /* 투자자/제작자 여부를 확인합니다. */

    /* USDC의 잔액을 확인합니다. */
    function balanceOfUSDC(address caller) public view returns (uint256) {
        return USDc.balanceOf(caller);
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

    function sendUSDC(
        address _to,
        address _from,
        uint256 $USDC
    ) public payable {
        // 잔량/송금액 확인
        require(balanceOfUSDC(_from) >= $USDC, "Caller don't have enough USDC");

        // Allowance 확인
        require(
            USDc.allowance(_from, _to) >= $USDC * 10**6,
            "Allowance is lower than caller claimed. Please check allowance."
        );

        USDc.transferFrom(_from, _to, $USDC * 10**6);
    }

    /* 해당 tokenId의 tokenURI를 관리하는 함수입니다. URIStorage라는 mapping 자료구조에 값을 넣어 저장합니다. */
    function _setURI(uint256 tokenId, string memory _tokenURI) internal {
        URIStorage[tokenId] = _tokenURI;
    }

    /* NFT minting, dest 주소로 amount개를 minting하고, tokenURI를 세팅합니다. */
    function mintNFT(address dest, string memory _tokenURI)
        internal
        returns (uint256)
    {
        // tokenID 증가
        _tokenIdCounter.increment();
        uint256 newTokenId = _tokenIdCounter.current();

        // mint
        _mint(dest, newTokenId);
        // setURI
        _setURI(newTokenId, _tokenURI);

        return newTokenId;
    }

    /* 영화를 민팅합니다. */
    function mintMovie(
        uint256 price,
        address makerAddress,
        string calldata _tokenURI
    ) external onlyOwner returns (uint256) {
        uint256 tokenId = mintNFT(owner(), _tokenURI);

        MovieTable[tokenId] = Movie(price, 0, makerAddress, _tokenURI);

        // mint
        emit MintMovie(tokenId, makerAddress);

        return tokenId;
    }

    /* 영화를 구매합니다. */
    function purchaseMovie(address buyer, uint256 tokenId) external onlyOwner {
        Movie memory movie = MovieTable[tokenId];

        /* 구매 기록이 없는지 확인합니다. */
        require(BuyerTable[tokenId][buyer] == Status.invalid);

        sendUSDC(address(this), buyer, movie.price);

        BuyerTable[tokenId][buyer] = Status.available;

        MovieTable[tokenId].totalRevenue += movie.price;
    }

    /* 투자금에 비례하여 수익 분배 비율을 계산합니다. */
    function getRevenueRatio(uint256 tokenId, address user)
        internal
        view
        returns (uint256)
    {
        address makerAddress = MovieTable[tokenId].maker;
        require(
            makerAddress == user || BuyerTable[tokenId][user] != Status.invalid,
            "Only maker or miner can access this function"
        );

        if (makerAddress == user) return MAKER_REVENUE_RATIO;

        return MINER_REVENUE_RATIO;
    }

    /* 투자금에 비례하여 수익 분배 비율을 계산합니다. */
    function calcRevenueRatio(
        uint256 tokenId,
        uint256 total,
        uint256 investment,
        address user
    ) internal view returns (uint256) {
        uint256 ratio = getRevenueRatio(tokenId, user);

        return ratio * (investment / total);
    }
}

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
