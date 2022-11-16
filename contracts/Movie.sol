// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/token/ERC1155/ERC1155Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC1155/extensions/ERC1155BurnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

contract Movie is
    Initializable,
    ERC1155Upgradeable,
    OwnableUpgradeable,
    ERC1155BurnableUpgradeable,
    UUPSUpgradeable
{
    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize() public initializer {
        __ERC1155_init("");
        __Ownable_init();
        __ERC1155Burnable_init();
        __UUPSUpgradeable_init();
    }

    function setURI(string memory newuri) public onlyOwner {
        _setURI(newuri);
    }

    function mint(
        address account,
        uint256 id,
        uint256 amount,
        bytes memory data
    ) public onlyOwner {
        _mint(account, id, amount, data);
    }

    function _authorizeUpgrade(address newImplementation)
        internal
        override
        onlyOwner
    {}

    /* 구매 여부를 확인합니다. */
    /* 투자자/제작자 여부를 확인합니다. */
    /* USDC 잔고를 확인합니다. */
    /* 투자자, 제작자로 등록합니다. */
    /* 거버넌스 토큰을 발급합니다. */
    /* 거버넌스 토큰 개수를 산정합니다. */
    /* 영화를 민팅합니다. */
    /* 영화를 구매합니다. */
    /* 투자자, 제작자 계좌에 판매 수익을 저장합니다. */
    /* 투자금에 비례하여 수익 분배 비율을 계산합니다. */
    /* 계좌에 에치된 USDC를 출금합니다. */
}
