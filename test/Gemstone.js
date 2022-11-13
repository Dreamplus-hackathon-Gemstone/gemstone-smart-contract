const {
  time,
  loadFixture,
} = require("@nomicfoundation/hardhat-network-helpers");
const { anyValue } = require("@nomicfoundation/hardhat-chai-matchers/withArgs");
const { expect } = require("chai");
const { before } = require("mocha");

describe("Gemstone", () => {
  before(() => {
    const [deployer] = await ethers.getSigners();
    const GemStone = await hre.ethers.getContractFactory("Gemstone");

    const gemStone = await GemStone.deploy();
    await gemStone.deployed();

    console.log(gemStone.address);
  })
});
