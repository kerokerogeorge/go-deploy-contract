// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract TestSmartContract {

	uint256 value;
	string mood;

	function store(uint256 number) public{
		value = number;
	}

	function retrieve() public view returns (uint256) {
		return value;
	}

  //create a function that writes a mood to the smart contract
  function setMood(string memory _mood) public{
      mood = _mood;
  }

  //create a function the reads the mood from the smart contract
  function getMood() public view returns(string memory){
      return mood;
  }
}
