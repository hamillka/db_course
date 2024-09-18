import React, { useState } from "react";
import { MenuItem } from "../../types/mytypes";

interface DropdownMenuProps {
  items: MenuItem[];
  onSelect: (item: MenuItem) => void;
}

const DropdownMenu: React.FC<DropdownMenuProps> = ({ items, onSelect }) => {
  const [selectedItem, setSelectedItem] = useState<{
    label: string;
    value: number;
  } | null>(null);

  const [open, setOpen] = useState(false);

  const handleItemClick = (item: MenuItem) => {
    setSelectedItem(item);
    onSelect(item);
    setOpen(false);
  };

  return (
    <div className="dropdown-menu">
      <button
        className="dropdown-button"
        style={{
          width: 150,
          marginTop: 10,
          fontSize: 15,
          borderTopLeftRadius: 10,
          borderTopRightRadius: 10,
          backgroundColor: "#f2f2ed",
          border: "1px solid #ddd",
        }}
        onClick={() => setOpen(!open)}
      >
        {selectedItem ? selectedItem.label : "Выбрать роль"}
      </button>
      {open && (
        <ul
          className="dropdown-list"
          style={{
            backgroundColor: "#f2f2ed",
            border: "1px solid #ddd",
            marginTop: 5,
            padding: 5,
            borderBottomLeftRadius: 10,
            borderBottomRightRadius: 10,
            width: 138,
            listStyle: "none",
          }}
        >
          {items.map((item) => (
            <li key={item.value} onClick={() => handleItemClick(item)}>
              {item.label}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default DropdownMenu;
