hat       = require 'hat'
jraphical = require 'jraphical'

{ argv }  = require 'optimist'
KONFIG    = require('koding-config-manager').load("main.#{argv.c}")


module.exports = class CredentialStore

  KodingError      = require '../../error'
  JCredentialData  = require './credentialdata'
  SocialCredential = require '../socialapi/credential'

  @SNEAKER_SUPPORTED = do ->

    for key, val of KONFIG.sneakerS3
      return no  if not val or val is ''

    return yes


  # STORE BEGINS --------------------------------------------------------------


  storeOnSneaker = (client, data, callback) ->

    { meta, identifier } = data

    _data = { pathName: identifier, data: meta }

    SocialCredential.store client, _data, (err) ->
      callback err, identifier


  storeOnMongo = (data, callback) ->

    credData = new JCredentialData data
    credData.save (err) ->
      callback err, data.identifier


  @create = (client, data, callback) ->

    { meta, originId, identifier } = data
    identifier ?= hat()
    data = { meta, originId, identifier }

    if @SNEAKER_SUPPORTED
      storeOnSneaker client, data, (err) ->
        # This part can be removed once kloud is ready
        # to use sneaker by default ~ GG cc/ RJ
        return callback err  if err
        storeOnMongo data, callback
    else
      storeOnMongo data, callback


  # STORE ENDS ----------------------------------------------------------------

