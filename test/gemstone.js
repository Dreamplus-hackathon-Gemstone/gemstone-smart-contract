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

async function deployBallot() {
	const Ballot = await hre.ethers.getContractFactory("Ballot");
	const [owner, addr1, addr2] = await hre.ethers.getSigners();
	const ballot = await upgrades.deployProxy(Ballot, []);
	await ballot.deployed();

	return { Ballot, ballot, owner, addr1, addr2 };
}

describe("TEST", () => {

	describe("Ballot Test", () => {

		it("Full coverage", async function () {
			const { Ballot, ballot, owner, addr1, addr2 } = await loadFixture(
				deployBallot
			);
			await ballot.connect(addr1).setApprovalForAll(ballot.address, true);
			await ballot.connect(addr2).setApprovalForAll(ballot.address, true);
			await ballot.connect(owner).setApprovalForAll(ballot.address, true);


			const contractOwner = await ballot.owner();
			expect(contractOwner, "Owner Test").to.equal(owner.address);

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

			await ballot.connect(owner).setAgenda(2, [0, 0, 0], "www.naver.com", addr2.address);
			const setAgendaTest = await ballot.connect(owner).showAgenda(1);
			expect(setAgendaTest.chairPerson).to.equal(addr2.address);


			await ballot.connect(owner).Register(1, 10, addr1.address);
			await ballot.connect(owner).Register(1, 20, owner.address);

			await ballot.connect(owner).Vote(1, 1, addr1.address);
			await ballot.connect(owner).Vote(1, 2, owner.address);
			const voteTest = await ballot.connect(owner).showAgenda(1);
			expect(voteTest.voters.length).to.equal(2);

			const ret = await ballot.connect(owner).Finish(1, addr2.address);

			const finishTest = await ballot.connect(owner).showAgenda(1);

			expect(finishTest.status).to.equal(2);
			expect(Number(finishTest.winner)).to.equal(2);
		})
	})

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
});


