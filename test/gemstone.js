const {
	time,
	loadFixture,
} = require("@nomicfoundation/hardhat-network-helpers");
const { anyValue } = require("@nomicfoundation/hardhat-chai-matchers/withArgs");
const { expect } = require("chai");
const { before, it } = require("mocha");
const { upgrades, ethers } = require("hardhat");
const hre = require("hardhat");
const { BigNumber } = require("ethers");
require("dotenv").config();

async function deployGemstone() {
	const GemStone = await hre.ethers.getContractFactory("GemstoneUpgradable");
	const [owner, addr1, addr2] = await hre.ethers.getSigners();
	const gemStone = await upgrades.deployProxy(GemStone, []);
	await gemStone.deployed();

	return { GemStone, gemStone, owner, addr1, addr2 };
}

async function deployUSDC() {
	const USDC = await hre.ethers.getContractFactory("USD");
	const [owner, addr1, addr2, addr3, addr4] = await hre.ethers.getSigners();
	const usdc = await USDC.deploy();

	await usdc.connect(owner).mint(addr1.address, BigNumber.from((10 ** 6) * 100000));
	await usdc.connect(owner).mint(addr2.address, BigNumber.from((10 ** 6) * 100000));
	await usdc.connect(owner).mint(addr3.address, BigNumber.from((10 ** 6) * 100000));
	await usdc.connect(owner).mint(addr4.address, BigNumber.from((10 ** 6) * 100000));
	return { USDC, usdc, owner, addr1, addr2, addr3, addr4 };
}


describe("Gemstone", () => {
	it("Contract Owner test", async function () {
		const { gemStone, owner } = await loadFixture(deployGemstone);
		const contractOwner = await gemStone.owner();
		expect(contractOwner).to.equal(owner.address);
	});

	it("Set Approve Test", async function () {
		const { Gemstone, gemStone, owner, addr1, addr2 } = await loadFixture(
			deployGemstone
		);
		await gemStone.connect(addr1).setApprovalForAll(gemStone.address, true);
		await gemStone.connect(addr2).setApprovalForAll(gemStone.address, true);
		await gemStone.connect(owner).setApprovalForAll(gemStone.address, true);

		const t1 = await gemStone
			.connect(owner)
			.isApprovedForAll(addr1.address, gemStone.address);
		const t2 = await gemStone
			.connect(addr1)
			.isApprovedForAll(addr1.address, gemStone.address);
		const t3 = await gemStone
			.connect(addr2)
			.isApprovedForAll(addr1.address, gemStone.address);

		expect(t1).to.equal(true);
		expect(t2).to.equal(true);
		expect(t3).to.equal(true);
	});
});