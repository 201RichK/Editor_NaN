import React, {Component} from 'react'
import './login.css'

export default class Login extends Component {
    constructor (props) {
        super(props)

        this.state = {
            username: '',
            password: '',
        }

        this.handleSubmit = this.handleSubmit.bind(this)
        this.handleChange = this.handleChange.bind(this)
    }

    handleChange(event){
        this.setState({
            username: event.target.username,
            password: event.target.password})
    }

    handleSubmit(event) {
        console.log(this.state.value)
       //event.preventDefault();
    }

    render = () => {
        return (
            <div className="vertical-align">
                <div className="container">
                    <div className="avatar"></div>
                    <form onSubmit={this.handleSubmit()}>

                        <input type="text" placeholder="Username" name="username" onChange={this.handleChange()}/>   
                        <input type="password" placeholder="Password" name="password" onChange={this.handleChange()}/>
                        <input type="submit" value="Login" />

                    </form>
                </div>
            </div>
        )
    }
}