var express = require('express');
var router = express.Router();
var log4js = require('log4js');
var logger = log4js.getLogger('SampleWebApp');
var invoke = require('./../help/invoke.js');
var query = require('./../help/query.js');


/* GET home page. */
router.get('/', function(req, res, next) {
	res.render('index', { 
		title: 'Bet Game',
		message:''
});
});
router.post('/', async function(req, res, next) {
	logger.debug('==================== INVOKE ON CHAINCODE ==================');
	var chaincodeName = 'bet';
	var channelName = 'mychannel';
	var fcn = 'newbet';
	var args = ['000000'];
	logger.debug('channelName  : ' + channelName);
	logger.debug('chaincodeName : ' + chaincodeName);
	logger.debug('fcn ttt : ' + fcn);
	logger.debug('args  : ' + args);
	if (!chaincodeName) {
		res.json(getErrorMessage('\'chaincodeName\''));
		return;
	}
	if (!channelName) {
		res.json(getErrorMessage('\'channelName\''));
		return;
	}
	if (!fcn) {
		res.json(getErrorMessage('\'fcn\''));
		return;
	}
	if (!args) {
		res.json(getErrorMessage('\'args\''));
		return;
	}
	let message = await invoke.invokeChaincode(/*peers, */channelName, chaincodeName, fcn, args/*, req.username, req.orgname*/);
	res.render('index',{
		title:'Bet Game',
		message:'New Game Is Begining',
	});
});

router.get('/guess',function(req,res,next){
  res.render('guess', {title:'Bet Game'});
});
router.post('/guess',async function(req,res,next){
  logger.debug('==================== INVOKE ON CHAINCODE ==================');
	var chaincodeName = 'bet';
	var channelName = 'mychannel';
	var fcn = 'creatbet';
	var args = [req.body.UserID,req.body.GuessNumber];
	logger.debug('channelName  : ' + channelName);
	logger.debug('chaincodeName : ' + chaincodeName);
	logger.debug('fcn ttt : ' + fcn);
	logger.debug('args  : ' + args);
	if (!chaincodeName) {
		res.json(getErrorMessage('\'chaincodeName\''));
		return;
	}
	if (!channelName) {
		res.json(getErrorMessage('\'channelName\''));
		return;
	}
	if (!fcn) {
		res.json(getErrorMessage('\'fcn\''));
		return;
	}
	if (!args) {
		res.json(getErrorMessage('\'args\''));
		return;
	}
	let message = await invoke.invokeChaincode(/*peers, */channelName, chaincodeName, fcn, args/*, req.username, req.orgname*/);
  res.render('index',{title:"Bet Game",message:'Your Number is submitted!'});

});


router.get('/check', function(req,res,next){
  res.render('check', {title:'Bet Game'});
});
router.post('/check',async function(req,res,next){
	var chaincodeName = 'bet';
	var channelName = 'mychannel';
	var fcn = 'readbet';
	var args = [req.body.UserID];
	var argssystem=['000000']
	logger.debug('channelName  : ' + channelName);
	logger.debug('chaincodeName : ' + chaincodeName);
	logger.debug('fcn  : ' + fcn);
	logger.debug('args  : ' + args);
	if (!chaincodeName) {
		res.json(getErrorMessage('\'chaincodeName\''));
		return;
	}
	if (!channelName) {
		res.json(getErrorMessage('\'channelName\''));
		return;
	}
	if (!fcn) {
		res.json(getErrorMessage('\'fcn\''));
		return;
	}
	if (!args) {
		res.json(getErrorMessage('\'args\''));
		return;
	}
	let messageone= await query.queryChaincode(/*peers, */channelName, chaincodeName, fcn, argssystem/*, req.username, req.orgname*/);
	System=JSON.parse(messageone)
	let message = await query.queryChaincode(/*peers, */channelName, chaincodeName, fcn, args/*, req.username, req.orgname*/);
	User=JSON.parse(message)
	var	Result='So sorry,Wrong!'
	if (System.Num==User.Num){
		res.render('congradulations', {
		});
	}
	res.render('result', {
		title:'Bet Game',
		SystemNumber:System.Num,
		UserNumber:User.Num,
		Result:Result,
	});
});

module.exports = router;
