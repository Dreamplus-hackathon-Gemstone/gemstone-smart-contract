const { upgrades } = require("hardhat");
const hre = require("hardhat");

async function main() {
	const [deployer] = await ethers.getSigners();
	const GemStone = await hre.ethers.getContractFactory("GemstoneUpgradable");
	console.log("Deploying Gemstone Upgradable contracts ... ");

	const gemStone = await upgrades.deployProxy(GemStone, []);
	await gemStone.deployed();
	console.log(`Deployer : ${gemStone.address}`);
	console.log(`Gemstone contract deployed at ${gemStone.address}`);
}

main().catch((error) => {
	console.error(error);
	process.exitCode = 1;
});
