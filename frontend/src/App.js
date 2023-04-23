import React, {useState} from 'react';
import logo from './pngfile/logo.png';
import textlogo from './pngfile/textlogo.png';
import texthistory from './pngfile/textHistory.png';
import './App.css';

function App() {
  
  const bgcolor = {
    backgroundColor: '#26225D',
    width : 'w-auto',
    height: 'h-auto',
    display : 'relative',
    top : '0px',
    left : '0px',
    
  }
  const sidebar = {
    backgroundColor: '#DC4E40',
    width: '275px',
    height: '673px',
    display : 'relative',
    borderRadius: '10px',
    padding: '20px',
    top : '0px',
    left : '-20px',
    position : 'relative'
  };
  
  const [selectedOption, setSelectedOption] = useState("");

  const handleOptionChange = (option) => {
    setSelectedOption(option);
  };
  
  const [text, setText] = useState('');

  const handleTextChange = (event) => {
    setText(event.target.value);
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
            name="option"
            value="Option 1"
            checked={selectedOption === "Option 1"}
            onChange={() => handleOptionChange("Option 1")}
          />
          <span className="radio-dot"></span> KMP
        </label>

        <label>
          <input
            type="radio"
            name="option"
            value="Option 2"
            checked={selectedOption === "Option 2"}
            onChange={() => handleOptionChange("Option 2")}
          />
          <span className="radio-dot"></span> BP
        </label>

      </div>
      <label className='textinput' htmlFor="text-input">
        <input
          type="text"
          id="text-input"
          value={text}
          onChange={handleTextChange}
        />
      {/* <p>You typed: {text}</p> */}
      </label>
    </div>
  );
}

export default App