import { ethers } from "hardhat";

async function main() {
  const socialPet = await ethers.deployContract("SocialPetExtended")
  await socialPet.waitForDeployment()

  console.log("socialPet: ", socialPet.accountAddress)
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
