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

async function deployBallot() {
	const Ballot = await hre.ethers.getContractFactory("Ballot");
	const [owner, addr1, addr2] = await hre.ethers.getSigners();
	const ballot = await upgrades.deployProxy(Ballot, []);
	await ballot.deployed();

	return { Ballot, ballot, owner, addr1, addr2 };
}


async function deployMovie(usdc) {
	const Movie = await hre.ethers.getContractFactory("MovieContract")
	const [owner, addr1, addr2, addr3, addr4] = await hre.ethers.getSigners();
	const movie = await upgrades.deployProxy(Movie, []);

	await movie.deployed();
	(await movie.connect(owner).setUSDC(usdc.address))
	return { Movie, movie, owner, addr1, addr2, addr3, addr4 };
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


	describe("Movie", function () {
		it("Deploy Test", async function () {
			const { USDC, usdc, owner, addr1, addr2, addr3, addr4 } = await loadFixture(deployUSDC);
			const { Movie, movie } = await deployMovie(usdc);

			await movie.connect(addr1).setApprovalForAll(movie.address, true);
			await movie.connect(addr2).setApprovalForAll(movie.address, true);
			await movie.connect(addr3).setApprovalForAll(movie.address, true);
			await movie.connect(addr4).setApprovalForAll(movie.address, true);
			await movie.connect(owner).setApprovalForAll(movie.address, true);
		})

		it("Mint Movie", async function () {
			const { USDC, usdc, owner, addr1, addr2, addr3, addr4 } = await loadFixture(deployUSDC);

			const { Movie, movie } = await deployMovie(usdc);

			await movie.connect(addr1).setApprovalForAll(movie.address, true);
			await movie.connect(addr2).setApprovalForAll(movie.address, true);
			await movie.connect(addr3).setApprovalForAll(movie.address, true);
			await movie.connect(addr4).setApprovalForAll(movie.address, true);
			await movie.connect(owner).setApprovalForAll(movie.address, true);

			await movie.connect(owner).mintMovie(20000, addr4.address, "www.naver.com");

			// const balanceOfAddr1 = await movie.connect(owner).balanceOfUSDC(addr1.address);
			// const balanceOfAddr2 = await movie.connect(owner).balanceOfUSDC(addr2.address);
			// const balanceOfAddr3 = await movie.connect(owner).balanceOfUSDC(addr3.address);
			// const balanceOfAddr4 = await movie.connect(owner).balanceOfUSDC(addr4.address);

			// expect(BigNumber.from(balanceOfAddr1)).to.deep.equal(BigNumber.from("100000"));
			// expect(BigNumber.from(balanceOfAddr2)).to.deep.equal(BigNumber.from("100000"));
			// expect(BigNumber.from(balanceOfAddr3)).to.deep.equal(BigNumber.from("100000"));
			// expect(BigNumber.from(balanceOfAddr4)).to.deep.equal(BigNumber.from("100000"));

			const res = await movie.connect(owner).getMovie(1);
			expect(res.price).to.deep.equal(BigNumber.from(20000));
			expect(res.maker).to.equal(addr4.address);
			expect(res.tokenURI).to.equal("www.naver.com");

		})

		it("Purchase Movie", async function () {
			const { USDC, usdc, owner, addr1, addr2, addr3, addr4 } = await loadFixture(deployUSDC);

			const { Movie, movie } = await deployMovie(usdc);

			await movie.connect(addr1).setApprovalForAll(movie.address, true);
			await movie.connect(addr2).setApprovalForAll(movie.address, true);
			await movie.connect(addr3).setApprovalForAll(movie.address, true);
			await movie.connect(addr4).setApprovalForAll(movie.address, true);
			await movie.connect(owner).setApprovalForAll(movie.address, true);

			await usdc.connect(addr1).approve(movie.address, BigNumber.from(1000000 * (10 ** 6)));
			await usdc.connect(addr2).approve(movie.address, BigNumber.from(1000000 * (10 ** 6)));
			await usdc.connect(addr3).approve(movie.address, BigNumber.from(1000000 * (10 ** 6)));
			await usdc.connect(addr4).approve(movie.address, BigNumber.from(1000000 * (10 ** 6)));
			await usdc.connect(owner).approve(movie.address, BigNumber.from(1000000 * (10 ** 6)));


			await movie.connect(owner).mintMovie(100, addr4.address, "www.naver.com");

			await movie.connect(owner).purchaseMovie(addr1.address, 1);
			await movie.connect(owner).purchaseMovie(addr2.address, 1);
			await movie.connect(owner).purchaseMovie(addr3.address, 1);

			const balanceOfAddr1 = await movie.connect(owner).balanceOfUSDC(addr1.address);
			const balanceOfAddr2 = await movie.connect(owner).balanceOfUSDC(addr2.address);
			const balanceOfAddr3 = await movie.connect(owner).balanceOfUSDC(addr3.address);
			const balanceOfContract = await usdc.connect(owner).balanceOf(movie.address);

			expect(BigNumber.from(balanceOfAddr1)).to.deep.equal(BigNumber.from(99900 * (10 ** 6)));
			expect(BigNumber.from(balanceOfAddr2)).to.deep.equal(BigNumber.from(99900 * (10 ** 6)));
			expect(BigNumber.from(balanceOfAddr3)).to.deep.equal(BigNumber.from(99900 * (10 ** 6)));
			expect(BigNumber.from(balanceOfContract)).to.deep.equal(BigNumber.from(300 * (10 ** 6)));
		})
	})
});


