let g:quickrun_config['go'] = {}
let g:quickrun_config['go']['command'] = 'make'
let g:quickrun_config['go']['cmdopt'] = 'go-test'
let g:quickrun_config['go']['exec'] = '%c %o SRC=@%'
let g:quickrun_config['go']['hook/cd/enable'] = 1
let g:quickrun_config['go']['hook/cd/directory'] = '$PWD'
