## Liquidstake upgrade

**Steps:**
1. Prepare proposal.

    **Important note**: If you want to set the `WhitelistAdminAddress`, you need to add the line `whitelist_admin_address: <your_whitelist_admin_address>` (e.g., tac15lvhklny0khnwy7hgrxsxut6t6ku2cgknw79fr) to the `plan.info` variable in the governance transaction message. The address must be valid; otherwise, an error log will appear during the update, and you will have to submit the governance transaction again.
2. Send gov tx to upgrade binary. (e.g., [gov_transaction.json](./gov_transaction.json))
3. Send admin tx to update module state.

    - Update params. (e.g., [1_admin_tx.json](./1_admin_tx.json))
    - Update validators. (e.g., [2_admin_tx.json](./2_admin_tx.json))

**Localnet Node Upgrade Checklist (v1.0.1 to v1.0.2)**

1. Initial Setup and Governance Proposal
    1. Checkout to the old node version.
    2. Build the old version.
        ```shell
        make build
        ```
    3. Initialize and start the local network (with a short voting period).
    ```shell
    make localnet-init GOV_TIME_SECONDS=60 TACCHAIND=./build/tacchaind
    make localnet-start TACCHAIND=./build/tacchaind
    ```
    4. (New Session) Submit the upgrade governance proposal (targeting height 400).
    ```shell
    ./build/tacchaind tx gov submit-proposal ./proposals/liquidstake_upgrade/gov_transaction.json --from validator --fees 200000000000000000utac --gas-adjustment 2 --gas 500000
    ```
    5. Vote 'Yes' on the proposal.
    ```shell
    ./build/tacchaind tx gov vote 1 yes --from validator --fees 80000000000000000utac
    ```
    6. Wait for the node to stop automatically at upgrade height (400).
2. Upgrade the Node Software
    1. Checkout to the new node version.
    2. Build the new version.
    ```shell
    make build
    ```
    3. Restart the node using the new binary.
    ```shell
    make localnet-start TACCHAIND=./build/tacchaind
    ```
    4. Verification: Node starts successfully at height 400.
3. Post-Upgrade Liquidstake Module Configuration
    1. Update Liquidstake module parameters (e.g., setting the admin address).
    ```shell
    tacchaind tx liquidstake update-params ./proposals/liquidstake_upgrade/1_admin_tx.json --from validator --fees 80000000000000000utac
    ```
    2. Unpause/enable the Liquidstake module.
    ```shell
    tacchaind tx liquidstake pause-module false --from validator --fees 80000000000000000utac
    ```

The upgraded node is successfully running, and the module is active.
