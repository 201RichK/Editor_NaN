# Editor_NaN





## Code de depart

### Code html5

```bash

#code

<div class="editor" v-on:paste.stop="doSomething">
     <editor editor-id="editor" :theme="theme" :langs="langs" :content="code" v-on:change-content="changeContentA"></editor>
</div>

# cdn

<script src="https://cdnjs.cloudflare.com/ajax/libs/vue/2.5.17/vue.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/ace/1.4.5/ace.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/ace/1.4.5/ext-language_tools.js"></script>
<script src="https://unpkg.com/axios/dist/axios.min.js"></script>

```

### Code Vue js
```bash

/* START: <ace-editor> Vue component */

Vue.component('Editor', {
    template: '<div :id="editorId" style="width: 100%; height: 100%;"></div>',
    props: ['editorId', 'content', 'langs', 'theme'],
    data() {
        return {
            editor: Object,
            beforeContent: '',

        }
    },
    watch: {
        'content' (value) {
            if (this.beforeContent !== value) {
                this.editor.setValue(value, 2)
            }
        }
    },
    mounted() {
        console.log(this.langs)
        var lang = this.langs || 'text'
        var theme = this.theme || 'github'
        window.ace.require("ace/ext/language_tools");
        this.editor = window.ace.edit(this.editorId)
        this.editor.setValue(this.content, 1)

        // mode et theme
        this.editor.getSession().setMode(`ace/mode/${lang}`)
        this.editor.setTheme(`ace/theme/${theme}`)
        this.editor.setOptions({
            highlightActiveLine: true,
            enableBasicAutocompletion: true,
            enableSnippets: true,
            enableLiveAutocompletion: true,
            fontSize: "13pt"
        });

        this.editor.on('change', () => {
            this.beforeContent = this.editor.getValue()
            this.$emit('change-content', this.editor.getValue())
        })

        this.editor.on('paste', function(Object, event) {
            return false
                //console.log('past')
        })

        //this.onload();
    },
    methods: {

    }
})


const app = new Vue({
    el: "#app",
    data() {
        return {

            result: '',

            theme: 'cobalt',
            langs: 'python',
            code: '',
            base_usl: 'http://127.0.0.1:8000',
            
        }
    },
    mounted: function() {

    },
    delimiters: ["${", "}"],

    methods: {
        reset() {
            this.code = 'reset content for Editor A'

        },
        changecode(val) {
            if (this.code !== val) {
                this.code = val
            }
        },
        update(edit) {
            console.log(edit)
            theme = edit
            console.log(theme)
            this.theme = edit
            this.editor = window.ace.edit('editor')

            this.editor.setTheme(`ace/theme/${theme}`)
        },

        sendcode(exoid) {
            this.editor = window.ace.edit('editor')

            console.log(exoid)

            code = this.editor.getValue()

            console.log(code)

            let data = JSON.stringify({
                code: code,

            })

            axios.defaults.xsrfCookieName = 'csrftoken'
            axios.defaults.xsrfHeaderName = 'X-CSRFToken'

            axios.post(this.base_usl + '/add', data, {
                    headers: {
                        'Content-Type': 'application/json',
                    }
                })
                .then((response) => {

                    resultats = JSON.parse(response.data)

                    console.log(resultats);

                })
                .catch((err) => {
                    console.log(err);
                })

        },

    }


})

```
