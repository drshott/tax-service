import React from "react";

// Function to format currency
export const formatCurrency = (amount) => {
    return new Intl.NumberFormat("en-IN", {
        style: "currency",
        currency: "INR",
        maximumFractionDigits: 0
    }).format(amount);
};

// Breakdown Section Component
const BreakdownSection = ({ breakdownData, title }) => {
    if (!breakdownData || !breakdownData.slabs) {
        return null; // Prevent rendering if no data
    }

    return (
        <div className="p-4 border rounded-lg shadow bg-white mt-4">
            {/* Title */}
            <h2 className="text-lg font-bold mb-3">{title}</h2>

            {/* Header Row */}
            <div className="grid grid-cols-4 text-sm font-semibold border-b pb-2 text-left">
                <span className="px-2">Slab</span>
                <span className="px-2">Rate</span>
                <span className="px-2">Amount</span>
                <span className="px-2">Tax</span>
            </div>

            {/* Breakdown Rows */}
            {breakdownData.slabs.map((item, index) => (
                <div key={index} className="grid grid-cols-4 text-sm py-2 border-b">
                    <span className="px-2">{item.range}</span>
                    <span className="px-2">{item.rate}</span>
                    <span className="px-2">{formatCurrency(item.amount)}</span>
                    <span className="px-2">{formatCurrency(item.slabTax)}</span>
                </div>
            ))}

            {/* Total Tax Row */}
            <div className="grid grid-cols-4 text-sm font-semibold border-t pt-2 mt-2">
                <span className="col-span-3 px-2">Tax</span>
                <span className="px-2">{formatCurrency(breakdownData.tax)}</span>
                <span className="col-span-3 px-2">Surcharge</span>
                <span className="px-2">{formatCurrency(breakdownData.surcharge)}</span>
                <span className="col-span-3 px-2">Cess</span>
                <span className="px-2">{formatCurrency(breakdownData.cess)}</span>
                <span className="col-span-3 px-2">Rebate</span>
                <span className="px-2">{formatCurrency(breakdownData.rebate)}</span>
                <span className="col-span-3 px-2">Prefessional Tax</span>
                <span className="px-2">{formatCurrency(2400)}</span>
                <span className="col-span-3 px-2">Total Tax</span>
                <span className="px-2">{formatCurrency(breakdownData.totalTax)}</span>
            </div>
        </div>
    );
};

export default BreakdownSection;
