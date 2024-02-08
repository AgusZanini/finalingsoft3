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
      const response = await axios.post('http://localhost:8080/operation', { number1: num1, number2: num2 });
      setResult(response.data);
    } catch (error) {
      console.error(error);
    }
  };

  const getOperationById = async () => {
    try {
      const response = await axios.get(`http://localhost:8080/operation/${id}`);
      setOperation(response.data);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div>
      <input type="number" value={num1} onChange={e => setNum1(parseFloat(e.target.value))} placeholder="Primer número" />
      <input type="number" value={num2} onChange={e => setNum2(parseFloat(e.target.value))} placeholder="Segundo número" />
      <button onClick={multiply}>Multiplicar</button>
      {result && <h3>Resultado: {result}</h3>}
      <input type="number" value={id} onChange={e => setId(e.target.value)} placeholder="ID de la operación" />
      <button onClick={getOperationById}>Obtener operación por ID</button>
      {operation && <h3>Operación: {JSON.stringify(operation)}</h3>}
    </div>
  );
}

export default App;
