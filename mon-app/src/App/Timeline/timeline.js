import React, {Component} from 'react'
import './timeline.css'

export default class Timeline extends Component {
    render = () => {
        return (
            <div className="timelineDiv flex justifyCenter">
                <div className="">
                    <div className="Timeline">
                        <svg height="5" width="200">
                            <line x1="0" y1="0" x2="200" y2="0" stroke="#004165" strokewith="5" />
                            Sorry, your browser does not support inline SVG.
                        </svg>

                        <div className="event1">

                            <div className="event1Bubble">
                            <div className="eventTime">
                                <div className="DayDigit">02</div>
                                <div className="Day">
                                    Run
                                    <div className="MonthYear">Exo 1</div>
                                </div>
                                <div className="DayDigit">02</div>
                                <div className="Day">
                                    Run
                                    <div className="MonthYear">Exo 1</div>
                                </div>
                            </div>
                            </div>
                            <div className="eventAuthor">John Doe succed with 4s runtime</div>
                            <svg height="20" width="20">
                            <circle cx="10" cy="11" r="5" fill="#004165" />
                            </svg>
                            <div className="time">4s runtime</div>

                        </div>

                        <svg height="5" width="300">
                            <line x1="0" y1="0" x2="300" y2="0" stroke="#004165" strokewith="5" />
                            Sorry, your browser does not support inline SVG.
                        </svg>

            <div className="event2">

                <div className="event2Bubble">
                    <div className="eventTime">
                        <div className="DayDigit">17</div>
                        <div className="Day">
                            Run
                            <div className="MonthYear">Exo2</div>
                        </div>
                        <div className="DayDigit">17</div>
                        <div className="Day">
                            Run
                            <div className="MonthYear">Exo2</div>
                        </div>
                    </div>
                </div>
                <div className="event2Author">john Doe Succed with 5s runtime</div>
                <svg height="20" width="20">
                    <circle cx="10" cy="11" r="5" fill="#004165" />
                </svg>
                <div className="time2">09s runtime</div>
            </div>

            <svg height="5" width="200">
                            <line x1="0" y1="0" x2="200" y2="0" stroke="#004165" strokewith="5" />
                            Sorry, your browser does not support inline SVG.
                        </svg>

                        <div className="event1">

                            <div className="event1Bubble">
                            <div className="eventTime">
                                <div className="DayDigit">02</div>
                                <div className="Day">
                                    Run
                                    <div className="MonthYear">Exo 1</div>
                                </div>
                                <div className="DayDigit">02</div>
                                <div className="Day">
                                    Run
                                    <div className="MonthYear">Exo 1</div>
                                </div>
                            </div>
                            </div>
                            <div className="eventAuthor">John Doe succed with 4s runtime</div>
                            <svg height="20" width="20">
                            <circle cx="10" cy="11" r="5" fill="#004165" />
                            </svg>
                            <div className="time">4s runtime</div>
                        </div>
                        <svg height="5" width="300">
                            <line x1="0" y1="0" x2="300" y2="0" stroke="#004165" strokewith="5" />
                            Sorry, your browser does not support inline SVG.
                        </svg>

            <div className="event2">

                <div className="event2Bubble">
                    <div className="eventTime">
                        <div className="DayDigit">17</div>
                        <div className="Day">
                            Run
                            <div className="MonthYear">Exo2</div>
                        </div>
                        <div className="DayDigit">17</div>
                        <div className="Day">
                            Run
                            <div className="MonthYear">Exo2</div>
                        </div>
                    </div>
                </div>
                <div className="event2Author">john Doe Succed with 5s runtime</div>
                <svg height="20" width="20">
                    <circle cx="10" cy="11" r="5" fill="#004165" />
                </svg>
                <div className="time2">09s runtime</div>
            </div>




            <svg height="5" width="50">
                <line x1="0" y1="0" x2="50" y2="0" stroke="#004165" strokewith="5" />
                Sorry, your browser does not support inline SVG.
            </svg>

            <div className="now">
                NOW
            </div>


            <svg height="5" width="150">
                <line x1="0" y1="0" x2="150" y2="0" stroke="rgba(162, 164, 163, 0.37)" strokewith="5" />
                Sorry, your browser does not support inline SVG.
            </svg>
            <div className="event3 futureGray ">
                <div className="event1Bubble futureOpacity">
                <div className="eventTime">
                    <div className="DayDigit">05</div>
                    <div className="Day">
                    Tuesday
                    <div className="MonthYear">May 2016</div>
                    </div>
                </div>
                <div className="eventTitle">Anticipated Hire</div>
                </div>
                <svg height="20" width="20">
                <circle cx="10" cy="11" r="5" fill="rgba(162, 164, 163, 0.37)" />
                </svg>
            </div>
            <svg height="5" width="50">
                <line x1="0" y1="0" x2="50" y2="0"  stroke="#004165" strokewith="5" />
            </svg>
            <svg height="20" width="42">
                <line x1="1" y1="0" x2="1" y2="20" stroke="#004165" strokewith="2" />
                <circle cx="11" cy="10" r="3" fill="#004165" />
                <circle cx="21" cy="10" r="3" fill="#004165" />
                <circle cx="31" cy="10" r="3" fill="#004165" />
                <line x1="41" y1="0" x2="41" y2="20" stroke="#004165" strokewith="2" />
            </svg>

            </div>
        </div>
    </div>
        )
    }
}