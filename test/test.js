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

describe("Gemstone", () => {

  async function deployGemstone() {
    const GemStone = await hre.ethers.getContractFactory("GemstoneUpgradable");
    const [owner, addr1, addr2] = await hre.ethers.getSigners();
    const gemStone = await upgrades.deployProxy(GemStone, []);
    await gemStone.deployed();

    return { GemStone, gemStone, owner, addr1, addr2 };
  }

  it("Contract Owner test", async function () {
    const { gemStone, owner } = await loadFixture(deployGemstone);
    const contractOwner = await gemStone.owner();
    expect(contractOwner).to.equal(owner.address);
  })

  it("Set Approve Test", async function () {
    const { Gemstone, gemStone, owner, addr1, addr2 } = await loadFixture(deployGemstone);
    const setApprove = await gemStone.connect(addr1).setApprovalForAll(gemStone.address, true);

    const isApproved = await gemStone.connect(addr1).isApprovedForAll(addr1.address, gemStone.address);
    expect(isApproved).to.equal(true);
  })
});
