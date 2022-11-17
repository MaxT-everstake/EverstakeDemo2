// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import EverstakeDemo2Dex from './everstake_demo2.dex'
import EverstakeDemo2Everstakedemo2 from './everstake_demo2.everstakedemo2'


export default { 
  EverstakeDemo2Dex: load(EverstakeDemo2Dex, 'everstake_demo2.dex'),
  EverstakeDemo2Everstakedemo2: load(EverstakeDemo2Everstakedemo2, 'everstake_demo2.everstakedemo2'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}