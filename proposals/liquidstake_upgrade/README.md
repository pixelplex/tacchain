## Liquidstake upgrade

**Steps:**
1. Prepare proposal.

    **Important note**: If you want to set the `WhitelistAdminAddress`, you need to add the line `whitelist_admin_address: <your_whitelist_admin_address>` (e.g., tac15lvhklny0khnwy7hgrxsxut6t6ku2cgknw79fr) to the `plan.info` variable in the governance transaction message. The address must be valid; otherwise, an error log will appear during the update, and you will have to submit the governance transaction again.
2. Send gov tx to upgrade binary. (e.g., [gov_transaction.json](./gov_transaction.json))
3. Send admin tx to update module state.

    - Update params. (e.g., [1_admin_tx.json](./1_admin_tx.json))
    - Update validators. (e.g., [2_admin_tx.json](./2_admin_tx.json))
