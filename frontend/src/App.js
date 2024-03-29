import React, {useState} from 'react';
import logo from './pngfile/logo.png';
import textlogo from './pngfile/textlogo.png';
import texthistory from './pngfile/textHistory.png';
import './App.css';
import axios from 'axios';

const botmessage = []

function App() {
  
  const bgcolor = {
    backgroundColor: '#26225D',
    height: '100%',
    width: '100%',
    display : 'flex',
    top : '0px',
    left : '0px',
    position : 'fixed',
    
  }
  const sidebar = {
    backgroundColor: '#DC4E40',
    width: '275px',
    height: '96vh',
    display : 'flex',
    borderRadius: '10px',
    padding: '20px',
    top : '0px',
    left : '-20px',
    position : 'fixed'
  };
  
  const [selectedOption, setSelectedOption] = useState("");

  const handleOptionChange = async (option) => {
    setSelectedOption(option);
    console.log(option);
    const algooption = await axios.post(`http://localhost:8000/api/algooption`, {"value":option});
    console.log(algooption);
  };
  
  const [text, setText] = useState('');

  const handleTextChange = (event) => {
    setText(event.target.value);
  };

  const [messages, setMessages] = useState([]);
  // const [botmessages, setBotMessages] = useState([]);

  const handleSend = async () => {
    if (text !== "") {
      setMessages([...messages, text]);
      setText("");
      const response = await axios.get(`http://localhost:8000/api/product/${text}`);
      const jaw = response.data.product;
      console.log(jaw);
      botmessage.push(jaw);
      // setBotMessages([...botmessages, { type: "bot", text: jaw }]); // add bot message
    }
  }

  const handleKeyDown = (event) => {
    if (event.key === 'Enter') {
      handleSend();
    }
  };

  return (
    <div style = {bgcolor}> 
      <div style = {sidebar}>
        
        <img src={logo} alt = "Logo" className="logo"/>
        <img src={textlogo} alt = "textLogo" className="textlogo" />
        <img src={texthistory} alt = "textHistory" className="texthistory"  />
        <button className="newchat">New Chat +</button>
        <button className="history" style={{top: "320px"}} ></button>
        <button className="history" style={{top: "390px"}} ></button>
        <button className="history" style={{top: "460px"}} ></button>
        <button className="history" style={{top: "530px"}} ></button>
        <button className="history" style={{top: "600px"}} ></button>

        <label>
          <input
            type="radio"
            name="option-kmp"
            value="KMP"
            checked={selectedOption === "Option 1"}
            onChange={() => handleOptionChange("Option 1")}
            style={{ display: 'none' }}
          />
          <div className="radio-style" style={{top: "700px", left: "100px"}}></div>
          <div className="KMP-text" style={{top:"713px", left: "150px"}}>KMP</div>
        </label>

        <label>
          <input
            type="radio"
            name="option"
            value="Option 2"
            checked={selectedOption === "Option 2"}
            onChange={() => handleOptionChange("Option 2")}
            style={{ display: 'none' }}
          />
          <span className="radio-style" style={{top: "730px",left: "73px"}}></span>
          <div className="BM-text" style={{top:"744px", left: "150px"}}>BM</div>
        </label>

      </div>
      <label className='textinput' htmlFor="text-input">
        <input
          className='input'
          type="text"
          id="text-input"
          value={text}
          onChange={handleTextChange}
          placeholder="Type your text here"
          onKeyDown={handleKeyDown}
          />
      {/* <p>You typed: {text}</p> */
          // console.log({text})
      }
      </label>
      <button className="send" style={{top: "740px"}} onClick={handleSend}> Send</button>
      
      <div className = "chat-container" style={{ overflowY: "scroll" }}>
      <div className="chatbox">
        
        {messages.map((message, index) => {
            return (
              <React.Fragment key={index}>
                <div className="chat-bubble chat-bubble-user">{message}</div>
                <div className="chat-bubble">{botmessage[index]}</div>
                
              </React.Fragment>
            );
        })}
      </div>
      </div>
    </div>
  );
}

export default App