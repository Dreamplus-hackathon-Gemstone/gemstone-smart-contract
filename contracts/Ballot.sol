// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/token/ERC1155/ERC1155Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC1155/extensions/ERC1155BurnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

contract Ballot is
    Initializable,
    ERC1155Upgradeable,
    OwnableUpgradeable,
    ERC1155BurnableUpgradeable,
    UUPSUpgradeable
{
    using Counters for Counters.Counter;

    Counters.Counter private _tokenIdCounter;

    /// @notice This event is emitted when Agenda opened.
    /// @param agendaId Token ID for Agenda
    /// @param chairPerson Address of chairePerson of Agenda
    event AgendaOpen(uint256 agendaId, address chairPerson);

    /// @notice This event is emitted when Agenda closed.
    /// @param agendaId Token ID for Agenda
    /// @param winner Winner value for Agenda
    event AgendaClose(uint256 agendaId, uint256 winner);

    /// @notice This event is emitted when Agenda removed.
    /// @param agendaId Token ID for Agenda
    /// @param chairPerson Winner value for Agenda
    event AgendaRemoved(uint256 agendaId, address chairPerson);

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _tokenIdCounter.increment();
        _disableInitializers();
    }

    enum Status {
        invalid,
        pending,
        done
    }

    struct Agenda {
        uint256 numberOfMiner;
        uint256 winner;
        string tokenURI;
        Status status;
        address chairPerson;
        uint256[] candidates;
        Voter[] voters;
    }

    struct Voter {
        uint256 agendaId;
        uint256 weight;
        uint256 voted;
        uint256 value;
        address voterAddress;
    }

    mapping(uint256 => Agenda) AgendaTable;

    mapping(uint256 => mapping(address => Voter)) VoterTable;

    /* 의장이 맞는지 확인합니다. */
    modifier isChairPerson(uint256 agendaId, address chairPerson) {
        // 의장이 맞는가?
        require(
            AgendaTable[agendaId].chairPerson == chairPerson,
            "Caller isn't chair person"
        );

        _;
    }

    /* Miner가 맞는지 확인합니다. */
    modifier isValidVoter(uint256 agendaId, address voterAddr) {
        // 스토리지 데이터 복사
        Voter memory voter = VoterTable[agendaId][voterAddr];

        require(voter.voterAddress == voterAddr, "Invalid Request");

        require(voter.voted == 0, "Already Voted");

        _;
    }

    function initialize() public initializer {
        __ERC1155_init("");
        __Ownable_init();
        __ERC1155Burnable_init();
        __UUPSUpgradeable_init();
    }

    function setURI(string memory newuri) public onlyOwner {
        _setURI(newuri);
    }

    function mint(
        address account,
        uint256 id,
        uint256 amount,
        bytes memory data
    ) public onlyOwner {
        _mint(account, id, amount, data);
    }

    function _authorizeUpgrade(address newImplementation)
        internal
        override
        onlyOwner
    {}

    /// @notice showAgenda
    /// @dev Explain to a developer any extra details
    /// @param agendaId Token ID for agenda
    /// @return Agenda Agenda Information
    function showAgenda(uint256 agendaId) external view returns(Agenda memory) {
        return AgendaTable[agendaId];
    }

    /// @notice setAgenda
    /// @param nom Number of miners.
    /// @param tokenURI Token URI for metadata of agenda.
    /// @param chairPerson Address of chairPerson.
    function setAgenda(
        uint256 nom,
        uint256[] calldata candidates,
        string calldata tokenURI,
        address chairPerson
    ) external onlyOwner {
        _tokenIdCounter.increment();
        uint256 agendaId = _tokenIdCounter.current();
        mint(chairPerson, agendaId, 1, "");

        AgendaTable[agendaId].numberOfMiner = nom;
        AgendaTable[agendaId].tokenURI = tokenURI;
        AgendaTable[agendaId].chairPerson = chairPerson;
        AgendaTable[agendaId].candidates = candidates;
        AgendaTable[agendaId].status = Status.pending;

        emit AgendaOpen(agendaId, chairPerson);
    }

    /// @notice BatchRegister
    /// @param agendaId Token ID for agenda
    /// @param weightArr Array of Weights of each voter
    /// @param addressArr Array of Addresses of each voter
    function BatchRegister(
        uint256 agendaId,
        uint256[] calldata weightArr,
        address[] calldata addressArr
    ) external onlyOwner {
        for (uint256 i = 0; i < weightArr.length; i++) {
            Register(agendaId, weightArr[i], addressArr[i]);
        }
    }

    /// @notice Register
    /// @param agendaId Token ID for agenda
    /// @param weight Weight of voter
    /// @param voterAddr Address of voter
    function Register(
        uint256 agendaId,
        uint256 weight,
        address voterAddr
    ) public onlyOwner {
        // VoterTable에 등록
        VoterTable[agendaId][voterAddr].agendaId = agendaId;
        VoterTable[agendaId][voterAddr].weight = weight;
        VoterTable[agendaId][voterAddr].value = 0;
        VoterTable[agendaId][voterAddr].voted = 0;
        VoterTable[agendaId][voterAddr].voterAddress = voterAddr;
    }

    /// @notice Vote
    /// @param agendaId Token ID for agenda
    /// @param value Candidate value for agenda
    /// @param voterAddr Address of voter
    function Vote(
        uint256 agendaId,
        uint256 value,
        address voterAddr
    ) external onlyOwner isValidVoter(agendaId, voterAddr) {
        require(
            value != 0 && value < AgendaTable[agendaId].candidates.length,
            "Invalid candidate"
        );
        // 스토리지 데이터 복사
        Voter memory voter = VoterTable[agendaId][voterAddr];
        // 데이터 갱신
        voter.value = value;
        voter.voted = 1;
        // 투표
        AgendaTable[agendaId].candidates[value] += voter.weight;
        AgendaTable[agendaId].voters.push(voter);
        // 투표 완료로 갱신
        VoterTable[agendaId][voterAddr].value = value;
        VoterTable[agendaId][voterAddr].voted = 1;
    }

    /// @notice Finish voting
    /// @param agendaId Token ID for Agenda
    /// @param chairPerson Address of chairPerson
    function Finish(uint256 agendaId, address chairPerson)
        external
        onlyOwner
        isChairPerson(agendaId, chairPerson)
    {
        Agenda memory agenda = AgendaTable[agendaId];

        // Pending 상태가 맞는지 확인.
        require(agenda.status == Status.pending, "Invalid Request");

        // 전체 Miner의 60%가 투표해야 종료 가능합니다.
        require(
            agenda.voters.length > agenda.numberOfMiner * 6 / 10,
            "Participation must be at least 60% of the number of whole miners"
        );

        // Winner value를 찾습니다.
        uint256 winner = CalcWinner(agendaId);

        // 상태 갱신
        AgendaTable[agendaId].status = Status(2);
        AgendaTable[agendaId].winner = winner;

        // Event
        emit AgendaClose(agendaId, winner);
    }

    /// @notice Check winner value
    /// @param agendaId Token ID for Agenda
    /// @return winner Result of voting
    function CalcWinner(uint256 agendaId)
        internal
        view
        returns (uint256 winner)
    {
        Agenda memory agenda = AgendaTable[agendaId];
        uint256 winningVoteCount = 0;
        for (uint256 i = 0; i < agenda.voters.length; i++) {
            Voter memory v = agenda.voters[i];
            agenda.candidates[v.value] += v.weight;
            if (agenda.candidates[v.value] > winningVoteCount) {
                winningVoteCount = agenda.candidates[v.value];
                winner = v.value;
            }
        }
        return winner;
    }

    /// @notice Remove Agenda
    /// @param agendaId Token ID for Agenda
    /// @param chairPerson Address of chairPerson
    function Remove(uint256 agendaId, address chairPerson)
        external
        onlyOwner
        isChairPerson(agendaId, chairPerson)
    {
        emit AgendaRemoved(agendaId, chairPerson);
    }
}
