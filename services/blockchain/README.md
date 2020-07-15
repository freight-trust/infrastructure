
## Rocks DB 
int DEFAULT_MAX_OPEN_FILES = 1024;
long DEFAULT_CACHE_CAPACITY = 8388608;
int DEFAULT_MAX_BACKGROUND_COMPACTIONS = 4;
int DEFAULT_BACKGROUND_THREAD_COUNT = 4;

MAX_OPEN_FILES_FLAG = "--Xplugin-rocksdb-max-open-files";
CACHE_CAPACITY_FLAG = "--Xplugin-rocksdb-cache-capacity";
MAX_BACKGROUND_COMPACTIONS_FLAG = "--Xplugin-rocksdb-max-background-compactions";
BACKGROUND_THREAD_COUNT_FLAG = "--Xplugin-rocksdb-background-thread-count";

## ETH Transactions
MAX_GET_HEADERS_FLAG = "--Xewp-max-get-headers";  "Maximum request limit for Ethereum Wire Protocol GET_BLOCK_HEADERS. (default: ${DEFAULT-VALUE})")
MAX_GET_BODIES_FLAG = "--Xewp-max-get-bodies"; "Maximum request limit for Ethereum Wire Protocol GET_BLOCK_BODIES. (default: ${DEFAULT-VALUE})"
MAX_GET_RECEIPTS_FLAG = "--Xewp-max-get-receipts";  "Maximum request limit for Ethereum Wire Protocol GET_RECEIPTS. (default: ${DEFAULT-VALUE})")
MAX_GET_NODE_DATA_FLAG = "--Xewp-max-get-node-data"; "Maximum request limit for Ethereum Wire Protocol GET_NODE_DATA. (default: ${DEFAULT-VALUE})")


MAX_GET_POOLED_TRANSACTIONS = "--Xewp-max-get-pooled-transactions"; "Maximum request limit for Ethereum Wire Protocol GET_POOLED_TRANSACTIONS. (default: ${DEFAULT-VALUE})")

ETH_65_ENABLED = "--Xeth-65-enabled";


{"--Xeip1559-basefee-max-change-denominator"}, basefeeMaxChangeDenominator = 8L;


{"--Xeip1559-target-gas-used"}, targetGasUsed = 10000000L;


{"--Xeip1559-slack-coefficient"}, slackCoefficient = 2L;


{"--Xeip1559-decay-range"}, decayRange = 1000000L;


{"--Xeip1559-initial-base-fee"}, initialBasefee = 1000000000L;


{"--Xeip1559-per-tx-gas-limit"}, perTxGasLimit = 8000000L;

{"--Xeip1559-basefee-max-change-denominator"}, basefeeMaxChangeDenominator = 8L;


{"--Xeip1559-target-gas-used"}, targetGasUsed = 10000000L;


{"--Xeip1559-slack-coefficient"}, slackCoefficient = 2L;


{"--Xeip1559-decay-range"}, decayRange = 1000000L;


{"--Xeip1559-initial-base-fee"}, initialBasefee = 1000000000L;


{"--Xeip1559-per-tx-gas-limit"}, perTxGasLimit = 8000000L;




.RLP

        names = "--start-time",
        description =
            "The timestamp in seconds of the first block for JSON imports. Subsequent blocks will be 1 second later. (default: current time)",
        arity = "1..1")
    private final Long startTime = System.currentTimeMillis() / 1000;