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

async function deployBallot() {
	const Ballot = await hre.ethers.getContractFactory("Ballot");
	const [owner, addr1, addr2] = await hre.ethers.getSigners();
	const ballot = await upgrades.deployProxy(Ballot, []);
	await ballot.deployed();

	return { Ballot, ballot, owner, addr1, addr2 };
}

describe("Ballot", () => {
	it("Contract Owner test", async function () {
		const { ballot, owner } = await loadFixture(deployBallot);
		const contractOwner = await ballot.owner();
		expect(contractOwner).to.equal(owner.address);
	});

	it("Set Approve Test", async function () {
		const { Ballot, ballot, owner, addr1, addr2 } = await loadFixture(
			deployBallot
		);
		await ballot.connect(addr1).setApprovalForAll(ballot.address, true);
		await ballot.connect(addr2).setApprovalForAll(ballot.address, true);
		await ballot.connect(owner).setApprovalForAll(ballot.address, true);

		const t1 = await ballot
			.connect(owner)
			.isApprovedForAll(addr1.address, ballot.address);
		const t2 = await ballot
			.connect(addr1)
			.isApprovedForAll(addr1.address, ballot.address);
		const t3 = await ballot
			.connect(addr2)
			.isApprovedForAll(addr1.address, ballot.address);

		expect(t1).to.equal(true);
		expect(t2).to.equal(true);
		expect(t3).to.equal(true);
	});
});
