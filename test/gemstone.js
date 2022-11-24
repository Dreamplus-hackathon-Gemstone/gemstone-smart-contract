const {
	time,
	loadFixture,
} = require("@nomicfoundation/hardhat-network-helpers");
const { anyValue } = require("@nomicfoundation/hardhat-chai-matchers/withArgs");
const { expect } = require("chai");
const { before, it } = require("mocha");
const { upgrades, ethers } = require("hardhat");
const hre = require("hardhat");
require("dotenv").config();

async function deployGemstone() {
	const GemStone = await hre.ethers.getContractFactory("GemstoneUpgradable");
	const [owner, addr1, addr2] = await hre.ethers.getSigners();
	const gemStone = await upgrades.deployProxy(GemStone, []);
	await gemStone.deployed();

	return { GemStone, gemStone, owner, addr1, addr2 };
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
