// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721";
import "@openzeppelin/contracts/access/Ownable";

contract SoulboundNFT is ERC721, Ownable {
    uint256 private _tokenId;
    string public baseURI;

    constructor(string memory _baseTokenURI) ERC721("Blockchain Credential SBT", "BCS") Ownable(msg.sender) {
        baseURI = _baseTokenURI;
        _tokenId = 0;
    }

    function mint(address to) external onlyOwner {
        _tokenId++;
        _safeMint(to, _tokenId);
    }

    function _update(address to, uint256 tokenId, address auth) internal override returns (address) {
        address from = _ownerOf(tokenId);
        require(from == address(0) || to == address(0), "Soulbound: non-transferable");
        return super._update(to, tokenId, auth);
    }

    function _baseURI() internal view override returns (string memory) {
        return baseURI;
    }
}
