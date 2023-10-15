import { expect } from "chai";
import { ethers } from "hardhat";
const {
  loadFixture
} = require("@nomicfoundation/hardhat-toolbox/network-helpers");

describe("SocialPet", function () {
  async function deploySocialPet() {
    const socialPet = await ethers.deployContract("SocialPetExtended")
    return { socialPet };
  }

  it("should mint a social pet nft", async function () {
    
    const socialPet = await loadFixture(deploySocialPet)
    const to = "0xC7Edcde376e29c25FfE1d539f9FE92c0E3E63E25"
    const tokenURI = "https://chocolate-objective-giraffe-337.mypinata.cloud/ipfs/Qmefgi96NSNCJFrUHWuPH67Y6kuk8KMayb9vZGgzVYLNtg?_gl=1*1isypq3*rs_ga*MTc4MjMzNjM0NC4xNjg0NTc2MTI4*rs_ga_5RMPXG14TE*MTY4NDU3NjEyOC4xLjEuMTY4NDU3NjE3MS4xNy4wLjA.";
    let amount = ethers.parseEther("0.1");
    const tx = await socialPet.mint(to, tokenURI, {
      value: amount
    })

    console.log("mint tx: ", tx.hash)
  });
});