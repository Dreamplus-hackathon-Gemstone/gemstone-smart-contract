const hre = require("hardhat");

async function getBalance(address) {
	const balanceBigInt = await hre.ethers.provider.getBalance(address);
	return hre.ethers.utils.formatEther(balanceBigInt);
}

async function main() {
	const [deployer] = await ethers.getSigners();
	const GemStone = await hre.ethers.getContractFactory("GEM");

	const gemStone = await GemStone.deploy();
	await gemStone.deployed();

	console.log(gemStone.address);
}

main().catch((error) => {
	console.error(error);
	process.exitCode = 1;
});
