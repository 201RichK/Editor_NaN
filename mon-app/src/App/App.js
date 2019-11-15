import React, {Component} from 'react'
import {BrowserRouter, Route} from 'react-router-dom'
import Login from './login/login'
import Playground from './playground/playground'

export default class App extends Component {
    render = () => {
        return(
            <BrowserRouter>
                <div className="main-route-place">
                    <Route exact path="/" component={Login} />
                    <Route path="/play" component={Playground} />
                </div>
            </BrowserRouter>
        )
    }
}