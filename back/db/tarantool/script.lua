#!/usr/bin/env tarantool

box.cfg {
    listen = 3301,
    log_format = 'plain',
    log_level = 5,
}

function init()
    s = box.schema.space.create('user', {if_not_exists=true})
    s:format({
        { name = 'user_id',  type = 'unsigned' },
        { name = 'cookie',   type = 'string' },
    })
    s:create_index('primary', { type = 'HASH', parts = {'cookie'}, if_not_exists = true })

    box.log.info('create space sessions')
end

function new_session(cookie, user_id)
    box.log.info('receive data for user %s', user_id)
    box.space.sessions:insert{cookie, user_id}
    box.log.info('insert cookie %s for %s user success', cookie, user_id)
end

function check_session(cookie)
    local session = box.space.sessions:select{cookie}
    box.log.info('found session for user %s', session.user_id)
    return session
end

function delete_session(cookie)
    local  droped_session = box.space.sessions:delete{cookie}
    box.log.info('delete session %s', deleted_session)
    return droped_session
end
