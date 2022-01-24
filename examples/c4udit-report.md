# c4udit Report

## Files analyzed
- examples/Test.sol

## Issues found

### Don't Initialize Variables with Default Value

Issue Information: [G001](https://github.com/byterocket/c4-common-issues/blob/main/0-Gas-Optimizations.md#g001---dont-initialize-variables-with-default-value)

Findings:
```
examples/Test.sol::4 => uint256 a = 0;
examples/Test.sol::12 => for (uint256 i = 0; i < array.length; i++) {
```

### Cache Array Length Outside of Loop

Issue Information: [G002](https://github.com/byterocket/c4-common-issues/blob/main/0-Gas-Optimizations.md#g002---cache-array-length-outside-of-loop)

Findings:
```
examples/Test.sol::12 => for (uint256 i = 0; i < array.length; i++) {
```

### Long Revert Strings

Issue Information: [G007](https://github.com/byterocket/c4-common-issues/blob/main/0-Gas-Optimizations.md#g007---long-revert-strings)

Findings:
```
examples/Test.sol::6 => string b = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
```

### Shift Right instead of Dividing by 2

Issue Information: [G008](https://github.com/byterocket/c4-common-issues/blob/main/0-Gas-Optimizations.md#g008---shift-right-instead-of-dividing-by-2)

Findings:
```
examples/Test.sol::13 => i = i / 2;
```

### Unsafe ERC20 Operation(s)

Issue Information: [L001](https://github.com/byterocket/c4-common-issues/blob/main/2-Low-Risk.md#l001---unsafe-erc20-operations)

Findings:
```
examples/Test.sol::16 => token.transferFrom(msg.sender, address(this), 100);
```

### Unspecific Compiler Version Pragma

Issue Information: [L003](https://github.com/byterocket/c4-common-issues/blob/main/2-Low-Risk.md#l003---unspecific-compiler-version-pragma)

Findings:
```
examples/Test.sol::1 => pragma solidity ^0.8.0;
```

