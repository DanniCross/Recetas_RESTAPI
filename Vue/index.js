    		const app = new Vue({
    			el: '#app',
    			data:{
    				pos:'',
    				receta: [],
                    crear:{
                        _id: '',
                        nombre: '',
                        ingredientes: [],
                        elaboracion: '',
                        pos: '',
                    },
                    update:{
                        _id: '',
                        nombre: '',
                        ingredientes: [],
                        elaboracion: '',
                        pos: '',
                    }
    			},
                methods:{
                    getCAPI: function(){
                        axios.get(`http://localhost:8000/recetas/${this.pos || ''}`).then(response => {
                            this.pos = '';
                            this.receta = response.data;
                        }).catch( e => {
                            console.log(e);
                        })
                    },
                    getAPI: function(){
                        axios.get(`http://localhost:8000/recetas/${this.pos || ''}`).then(response => {
                            if(this.pos != ''){
                                this.pos = '';
                                this.receta = response.data;
                            }
                        }).catch(e => {
                            console.log(e);
                        })
                    },
                    setAPI: function(){
                        this.crear.pos = this.pos;
                        axios.post(`http://localhost:8000/recetas/${this.pos || ''}`, this.crear).then(response => {
                            this.pos = '';
                            this.crear.nombre = '';
                            this.crear.ingredientes = [];
                            this.crear.elaboracion = '';
                            this.receta = response.data;
                        }).catch(e => {
                            console.log(e);
                        })
                    },
                    updateAPI: async function(){
                            const response = await axios.put(`http://localhost:8000/recetas/`,this.update);
                            this.getAPI();
                            this.getReceta();
                            this.receta = response.data;
                            //return response.data;
                    },
                    getReceta: async function() {
                        const response = await axios.get(`http://localhost:8000/recetas/${this.pos || ''}`);
                        this.recetas = response.data;
                        return response.data;
                    },
                    edit: function(receta){
                        for(const key in receta){
                            if(receta.hasOwnProperty(key)){
                                this.update[key] = receta[key];
                            }
                        }
                    },
                    deleteAPI: function(){
                        axios.delete(`http://localhost:8000/recetas/${this.pos}`).then(response => {
                            this.receta = response.data;
                        }).catch(e => {
                            console.log(e);
                        })
                    },
                }
    		});