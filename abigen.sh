$(which solc) --abi contracts/Gemstone.sol --include-path node_modules --base-path . -o build --overwrite
$(which abigen) --abi build/GEM.abi --pkg main --type GEM --out GEM.go