const { task, vars } = require("hardhat/config");
const { getVaultInstance } = require("./common");

/* run this command before running any tasks below:
npx hardhat vars set FUNDINGVAULT_USER_PRIVATE_KEY
*/

task("claim", "Claim an amount in wei from a specified ID or any grant(s) the sender wallet holds")
  .addParam("amount", "(wei) Set 0 to claim all available funds")
  .addOptionalParam("id", "Optional: Specify Grant ID to claim")
  .addOptionalParam("target", "Optional: Sends funds to the target address. Funds will be sent to the sender if left blank")
  .setAction(async (args) => {
    const userWallet = getUserWallet();
    const vaultContract = getVaultInstance();

    let tx;
    if (args.id) {
        if (args.target) {
            tx = await vaultContract.connect(userWallet).claimTo(
                args.id,
                args.amount,
                args.target
            );
        } else {
            tx = await vaultContract.connect(userWallet).claim(
                args.id,
                args.amount
            );
        }
    } else {
        if (args.target) {
            tx = await vaultContract.connect(userWallet).claimTo(
                args.amount,
                args.target
            );
        } else {
            tx = await vaultContract.connect(userWallet).claim(
                args.amount,
                args.target
            );
        }
    }

    console.log("Successful claim. Tx: ", tx.hash);
  })


function getUserWallet() {
    let userKey = vars.get('FUNDINGVAULT_USER_PRIVATE_KEY');
    let userWallet = new ethers.Wallet(userKey, ethers.provider);
    return userWallet;
}