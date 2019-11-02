import axios from 'axios'

export default {

    baseURI: 'http://localhost:8000/api',
    //  baseURI: '/api',

    // APIs for search object
    getItemTemplates(cb) {
        //axios.get(this.baseURI + '/itemTemplates').then(r => cb(r.data))
        return [{
                "id": "bow_001",
                "name": "Basic Bow",
                "description": "A bow",
                "itemType": 3,
                "properties": {
                    "range": 15,
                    "damage": 1.5
                }
            },
            {
                "id": "sword_001",
                "name": "Basic Sword",
                "description": "A sword like any other",
                "itemType": 2,
                "properties": {
                    "range": 2,
                    "damage": 1.2
                }
            }
        ]
    },
    deleteSearchByID(id, cb) {
        axios.defaults.withCredentials = false;
        axios.delete(this.baseURI + '/itemTemplates/' + id).then(r => cb(r.data))
    },
    createItemTemplate(itemTemplate, cb, errorCb) {
        axios.post(this.baseURI + '/itemTemplates', itemTemplate)
            .then(r => {
                console.log(r);
                cb(r);
            })
            .catch(error => {
                console.log(error);
                errorCb(error);
            })
    }
}