var acc = require("./acc_pb")
var mainID = require("./mainID_pb")
var errorcode = require("./errorcode_pb")
var logic = require("./logic_pb")
var server = require("./server_pb")

module.exports={
    AccPb :     acc,
    MainPb:     mainID,
    ErrorCode:  errorcode,
    LogicPb:    logic,
    ServerPb:   server,
}