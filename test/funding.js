const {
	time,
	loadFixture,
} = require('@nomicfoundation/hardhat-network-helpers');
const { anyValue } = require('@nomicfoundation/hardhat-chai-matchers/withArgs');
const { expect } = require('chai');
const { before, it } = require('mocha');
const { upgrades, ethers } = require('hardhat');
const hre = require('hardhat');
const { BigNumber } = require('ethers');
require('dotenv').config();

async function deployUSDC() {
	const USDC = await hre.ethers.getContractFactory('USD');
	const [owner, addr1, addr2, addr3, addr4] = await hre.ethers.getSigners();
	const usdc = await USDC.deploy();

	await usdc
		.connect(owner)
		.mint(addr1.address, BigNumber.from(10 ** 6 * 100000));
	await usdc
		.connect(owner)
		.mint(addr2.address, BigNumber.from(10 ** 6 * 100000));
	await usdc
		.connect(owner)
		.mint(addr3.address, BigNumber.from(10 ** 6 * 100000));
	await usdc
		.connect(owner)
		.mint(addr4.address, BigNumber.from(10 ** 6 * 100000));
	return { USDC, usdc, owner, addr1, addr2, addr3, addr4 };
}

async function deployFunding() {
	const Funding = await hre.ethers.getContractFactory('Funding');
	const [owner, addr1, addr2, addr3, addr4] = await hre.ethers.getSigners();
	const funding = await upgrades.deployProxy(Funding, []);
	await funding.deployed();

	await funding.connect(addr1).setApprovalForAll(funding.address, true);
	await funding.connect(addr2).setApprovalForAll(funding.address, true);
	await funding.connect(addr3).setApprovalForAll(funding.address, true);
	await funding.connect(addr4).setApprovalForAll(funding.address, true);

	return { Funding, funding, owner, addr1, addr2, addr3, addr4 };
}

describe('Funding', () => {
	it('Deploy', async function () {
		const { Funding, funding, owner, addr1, addr2, addr3, addr4 } =
			await loadFixture(deployFunding);

		const ownerAddress = await funding.connect(owner).owner();
		expect(owner.address).to.equal(ownerAddress);
	});

	it('Propose', async function () {
		const { Funding, funding, owner, addr1, addr2, addr3, addr4 } =
			await loadFixture(deployFunding);

		await funding
			.connect(addr1)
			.propose(10000 * 10 ** 6, 100, 'www.google.com');
		const view = await funding.connect(addr1).viewProposal(1);

		expect(view.makerAddress).to.equal(addr1.address);
		expect(view.tokenId).to.deep.equal(BigNumber.from(1));
		expect(view.targetAmount).to.deep.equal(BigNumber.from(10000 * 10 ** 6));
	});

	it('Invest', async function () {
		const { Funding, funding, owner, addr1, addr2, addr3, addr4 } =
			await loadFixture(deployFunding);
	});

	it('Withdraw', async function () {
		const { Funding, funding, owner, addr1, addr2, addr3, addr4 } =
			await loadFixture(deployFunding);
	});

	it('Funding Success', async function () {});

	it('Burn Proposal', async function () {});

	it('Exceptions', async function () {});
});
