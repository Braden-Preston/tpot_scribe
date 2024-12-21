import Quill from 'quill'
import './app.css'
// import './quill.core.css'
// import './quill.snow.css'

import { LoadDocument } from '../wailsjs/go/main/App'

async function main() {

    let html = await LoadDocument('test/paragraph.docx')
    document.querySelector('#editor').innerHTML = `
        <h1>Heading 1</h1>
        <h2>Heading 2</h2>
        <h3>Heading 3</h3>
        <h4>Heading 4</h4>
        <h5>Heading 5</h5>
        <h6>Heading 6</h6>
        This is normal sized text
        <ul>
            <li>Ordered A</li>
            <li>Ordered B</li>
        </ul>
        <ol>
            <li>Numbered A</li>
            <li>Numbered B</li>
        </ol>
        <p>This is <span style="color: #800000;">colored</span> text. <span style="text-decoration: underline;"><strong><em>And</em></strong></span> some <span style="background: #ffff00;">highlights</span>.</p>
        <p>This is a quote</p>
        <p>And breakpoint:</p>
        <hr />
        And a quote:
        <blockquote>This is a quote</blockquote>
        Remaining text.
    `

    let containerEl = document.querySelector('#editor-container')
    let toolbarEl = document.querySelector('#toolbar')
    let editorEl = document.querySelector('#editor')

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
        'hr',
        'pagebreak'
    ]

    const options = {
        theme: 'snow',
        placeholder: 'To get started, begin typing here...',
        formats: validFormats,
        modules: {
            toolbar: toolbarEl
        }
    }

    const quill = new Quill(editorEl, options)


}

main()
