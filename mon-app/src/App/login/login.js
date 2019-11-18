import React, {Component} from 'react'
import './login.css'
import * as yup from 'yup'
import Axios from 'axios'

export default class Login extends Component {
    constructor () {
        super()
        this.state = {
            email: '',
            password: '',
            emailError: '',
            passwordError: '',
        }
    }


    onChangeEmail = (event) => {
        this.setState({
            email: event.target.value
        })
    }
    onChangePassword = (event) => {
        this.setState({
            password: event.target.value
        })
    }

    onSubmit = (event) => {
        event.preventDefault()
        this.setState({
            emailError: '',
            passwordError: ''
        })
        let schema = yup.object().shape({
            password: yup.string().required().min(8, "password must be superior 8 character"),
            email: yup.string().email().required(),
          });
          
          schema.validate({
            email: this.state.email,
            password: this.state.password
          }).then((data) => {
              console.log(data)
            Axios.post("http://localhost:8080/login", data)
          }).catch((err) => {
              switch (err.path) {
                  case "email":
                      this.setState({
                          emailError: err.errors[0]
                      })
                      break
                case "password":
                    this.setState({
                        passwordError: err.errors[0]
                    })
              }
          })
        
    }

    onChange(event) {

    }

    render = () => {
        return (
            <div className="vertical-align">
                <div className="container">
                    <div className="avatar"></div>
                    <form onSubmit={this.onSubmit}>

                        <input type="text" placeholder="Username" name="username" onChange={this.onChangeEmail} />
                        { this.state.emailError !== ''  && this.state.emailError}   
                        <input type="password" placeholder="Password" name="password" onChange={this.onChangePassword} />
                        { this.state.passwordError !== ''  && this.state.passwordError} 
                        <input type="submit" value="Login" />

                    </form>
                </div>
            </div>
        )
    }
}