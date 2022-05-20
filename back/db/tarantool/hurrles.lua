#!/usr/bin/env tarantool

box.cfg {
    listen = 3301,
    log_format = 'plain',
    log_level = 5,
}

local log = require('log')

function init()
    s = box.schema.space.create('sessions', {if_not_exists=true})
    s:format({
        { name = 'cookie',   type = 'string' },
        { name = 'user_id',  type = 'unsigned' },
    })
    s:create_index('primary', { type = 'HASH', parts = {'cookie'}, if_not_exists = true })

    log.info('create space sessions')
end

function new_session(cookie, user_id)
    log.info('receive data for user %s', user_id)
    local session = box.space.sessions:insert{cookie, user_id}
    log.info('insert cookie %s for %s user success', cookie, user_id)
    return session
end

function check_session(cookie)
    local session = box.space.sessions:select{cookie}
    log.info('found session for user %s', session.user_id)
    return session
end

function delete_session(cookie)
    local  droped_session = box.space.sessions:delete{cookie}
    log.info('delete session %s', deleted_session)
    return droped_session
end
