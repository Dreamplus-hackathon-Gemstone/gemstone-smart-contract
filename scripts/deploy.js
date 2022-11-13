const hre = require("hardhat");

async function getBalance(address) {
  const balanceBigInt = await hre.ethers.provider.getBalance(address);
  return hre.ethers.utils.formatEther(balanceBigInt);
}

async function main() {
  const [deployer] = await ethers.getSigners();
  const GemStone = await hre.ethers.getContractFactory("Gemstone");

  const gemStone = await GemStone.deploy();
  await gemStone.deployed();

  console.log(`Deployer : ${gemStone.address}`);
  console.log(`Gemstone contract deployed at ${gemStone.address}`);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
