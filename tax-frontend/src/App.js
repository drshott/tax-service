import React, { useState } from 'react';
import BreakdownSection, {formatCurrency} from './breakdown';

const App = () => {
  const [income, setIncome] = useState('');
  const [tax25, setTax25] = useState(null);
  const [tax24, setTax24] = useState(null);
  const [taxNew, setTaxNew] = useState(null);
  const [taxOld, setTaxOld] = useState(null);
  const [error, setError] = useState(null);
  const [formData, setFormData] = useState({});

  const handleIncomeChange = (e) => {
    setIncome(e.target.value);
  };

  const calculateTax = async () => {

    //e.preventDefault();

    //setIncome(e.target.value);

    if (!income) {
      setError('Please enter an income amount');
      return;
    }
    try {
      // Replace with your external API URL for tax calculation
      //const response1 = await fetch("/api/tax", {
      //  method: "POST",
      //  headers: { "Content-Type": "application/json" },
      //  body: JSON.stringify(formData),
      //});
      const response = await fetch(`/api/tax/?income=${income}`);
      const data = await response.json();
      
      if (response.ok) {
        setTax25(data.taxNew.totalTax); // Assuming the API returns the tax value
        setTax24(data.taxOld.totalTax); // Assuming the API returns the
        setTaxNew(data.taxNew); // Assuming the API returns the tax value
        setTaxOld(data.taxOld);
        setError(null);
      } else {
        setError('Error calculating tax');
      }
    } catch (err) {
      setError('Failed to fetch tax data');
    }
  };

  return (
    <div class="max-w-4xl mx-auto p-4 space-y-4">
        <div class="bg-white rounded-lg shadow-md p-6">
            <h1 class="text-2xl font-bold mb-4">Compare 2024 and 2025 New Tax Regimes</h1>
            
            <div class="space-y-4">
                <div>
                    <label class="block text-sm font-medium mb-2">Enter Annual Income (₹)</label>
                    <input 
                        type="text" 
                        value={income}
                        placeholder="Enter amount"
                        onChange={handleIncomeChange}
                        class="w-full p-2 border rounded-md"
                    />
                </div>
                <button 
                        class="w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700"
						            onClick={calculateTax}
                >
                        Show Details
                </button>

                <div id="results">
                    <div class="grid md:grid-cols-2 gap-4">
                        <div class="bg-blue-50 p-4 rounded-lg">
                            <h3 class="font-semibold mb-2">Tax as per 2024 - New Regime</h3>
                            {tax24 !== null && !error && <p id="oldTax" class="text-2xl font-bold text-blue-600">{formatCurrency(tax24)}</p>}
                        </div>
                        <div class="bg-green-50 p-4 rounded-lg">
                            <h3 class="font-semibold mb-2">Tax as per 2025 - New Regime</h3>
                            {tax25 !== null && !error && <p id="newTax" class="text-2xl font-bold text-green-600">{formatCurrency(tax25)}</p>}
                        </div>
                    </div>
                    {taxOld && <h2 className="text-lg font-bold mb-3">Breakdown 2024 New Regime</h2>}
                    {taxOld && <BreakdownSection breakdownData={taxOld} />}
                    {taxNew && <h2 className="text-lg font-bold mb-3">Breakdown 2025 New Regime</h2>}
                    {taxNew && <BreakdownSection breakdownData={taxNew} />}
                </div>
            </div>
        </div>
      {error && <p style={{ color: 'red' }}>{error}</p>}
    </div>
  );
};

export default App;
