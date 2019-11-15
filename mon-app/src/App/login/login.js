import React, {Component} from 'react'
import './login.css'

export default class Login extends Component {
    render = () => {
        return (
            <div className="vertical-align">
                <div className="container">
                    <div className="avatar"></div>
                    <form method="POST">
                        <input type="text" placeholder="Username" name="username" />   
                        <input type="password" placeholder="Password" name="password"/>
                        <input type="submit" value="Login" />
                    </form>
                </div>
            </div>
        )
    }
}