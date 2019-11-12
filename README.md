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





This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

## Available Scripts

In the project directory, you can run:

### `npm start`

Runs the app in the development mode.<br />
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.<br />
You will also see any lint errors in the console.

### `npm test`

Launches the test runner in the interactive watch mode.<br />
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

### `npm run build`

Builds the app for production to the `build` folder.<br />
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.<br />
Your app is ready to be deployed!

See the section about [deployment](https://facebook.github.io/create-react-app/docs/deployment) for more information.

### `npm run eject`

**Note: this is a one-way operation. Once you `eject`, you can’t go back!**

If you aren’t satisfied with the build tool and configuration choices, you can `eject` at any time. This command will remove the single build dependency from your project.

Instead, it will copy all the configuration files and the transitive dependencies (Webpack, Babel, ESLint, etc) right into your project so you have full control over them. All of the commands except `eject` will still work, but they will point to the copied scripts so you can tweak them. At this point you’re on your own.

You don’t have to ever use `eject`. The curated feature set is suitable for small and middle deployments, and you shouldn’t feel obligated to use this feature. However we understand that this tool wouldn’t be useful if you couldn’t customize it when you are ready for it.

## Learn More

You can learn more in the [Create React App documentation](https://facebook.github.io/create-react-app/docs/getting-started).

To learn React, check out the [React documentation](https://reactjs.org/).

### Code Splitting

This section has moved here: https://facebook.github.io/create-react-app/docs/code-splitting

### Analyzing the Bundle Size

This section has moved here: https://facebook.github.io/create-react-app/docs/analyzing-the-bundle-size

### Making a Progressive Web App

This section has moved here: https://facebook.github.io/create-react-app/docs/making-a-progressive-web-app

### Advanced Configuration

This section has moved here: https://facebook.github.io/create-react-app/docs/advanced-configuration

### Deployment

This section has moved here: https://facebook.github.io/create-react-app/docs/deployment

### `npm run build` fails to minify

This section has moved here: https://facebook.github.io/create-react-app/docs/troubleshooting#npm-run-build-fails-to-minify
