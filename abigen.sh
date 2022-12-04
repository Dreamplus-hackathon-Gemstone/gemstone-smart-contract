$(which solc) --abi contracts/Ballot.sol --include-path node_modules --base-path . -o build --overwrite
$(which abigen) --abi build/Ballot.abi --pkg main --type GEM --out Ballot.go