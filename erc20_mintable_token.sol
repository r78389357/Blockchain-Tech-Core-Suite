// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20";
import "@openzeppelin/contracts/access/Ownable";
contract MintableERC20Token is ERC20, Ownable {
    string public constant TOKEN_NAME = "GlobalChain Token";
    string public constant TOKEN_SYMBOL = "GCT";
    uint8 public constant TOKEN_DECIMALS = 18;

    constructor() ERC20(TOKEN_NAME, TOKEN_SYMBOL) Ownable(msg.sender) {}

    function mint(address to, uint256 amount) external onlyOwner {
        _mint(to, amount * (10 ** uint256(TOKEN_DECIMALS)));
    }

    function burn(uint256 amount) external {
        _burn(msg.sender, amount);
    }

    function burnFrom(address account, uint256 amount) external {
        _spendAllowance(account, msg.sender, amount);
        _burn(account, amount);
    }
}
