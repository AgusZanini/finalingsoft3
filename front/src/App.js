import React, { useState } from 'react';
import axios from 'axios';

function App() {
  const [num1, setNum1] = useState("");
  const [num2, setNum2] = useState("");
  const [result, setResult] = useState("");
  const [id, setId] = useState("");
  const [operation, setOperation] = useState(null);

  const multiply = async () => {
    try {
      const response = await axios.post('https://backend-fuz7myfmqq-uc.a.run.app/operation', { number1: num1, number2: num2 });
      setResult(response.data);
    } catch (error) {
      console.error(error);
    }
  };

  const getOperationById = async () => {
    try {
      const response = await axios.get(`https://backend-fuz7myfmqq-uc.a.run.app/operation/${id}`);
      setOperation(response.data);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div>
      <input id="num1" type="number" value={num1} onChange={e => setNum1(parseFloat(e.target.value))} placeholder="Primer número" />
      <input id="num2" type="number" value={num2} onChange={e => setNum2(parseFloat(e.target.value))} placeholder="Segundo número" />
      <button id="multiply" onClick={multiply}>Multiplicar</button>
      {result && <h3 id="result">Resultado: {result}</h3>}
      <input id="idinsert" type="number" value={id} onChange={e => setId(e.target.value)} placeholder="ID de la operación" />
      <button id="getbyid" onClick={getOperationById}>Obtener operación por ID</button>
      {operation && <h3 id="operation">Operación: {JSON.stringify(operation)}</h3>}
    </div>
  );
}

export default App;
