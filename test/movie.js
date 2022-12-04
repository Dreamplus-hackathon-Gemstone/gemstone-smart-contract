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


	it("Get Revenue Ratio Test", async function () {
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

		const minerRatio = await movie.connect(addr1).getRevenueRatio(1, addr1.address);
		const makerRatio = await movie.connect(addr4).getRevenueRatio(1, addr4.address);

		expect(makerRatio).to.deep.equal(BigNumber.from(50))
		expect(minerRatio).to.deep.equal(BigNumber.from(40))
	})
})