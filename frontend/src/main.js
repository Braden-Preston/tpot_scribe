import Quill from 'quill'
import './app.css';
// import './quill.core.css'
import './quill.snow.css'

import { Greet } from "../wailsjs/go/main/App"

// alert("test")

// let msg = await Greet()

// document.querySelector("#editor").innerHTML = msg

const validFormats = [
    'background',
    'bold',
    'color',
    'italic',
    'link',
    'underline',
    'blockquote',
    'header',
    'list',
    'align',
    'image',
    'code',
    'divider'
]


const options = {
    theme: 'snow',
    placeholder: "To get started, begin typing here...",
    formats: validFormats,
    modules: {
        toolbar: {

        },
    },
}

const quill = new Quill('#editor', options);


