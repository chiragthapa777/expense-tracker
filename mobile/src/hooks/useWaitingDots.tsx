import { useEffect, useState } from "react";

export const useWaitingDots = () => {
  const [waitingDots, setWaitingDots] = useState("   ");

  useEffect(() => {
    const interval = setInterval(() => {
      setWaitingDots((prevDots) => {
        // Remove spaces and count dots
        const dotCount = prevDots.replace(/\s/g, "").length;

        if (dotCount === 3) {
          return "   "; // Reset to 3 spaces
        }

        // Add a dot and pad with spaces to maintain width of 3
        const newDotCount = dotCount + 1;
        return ".".repeat(newDotCount) + " ".repeat(3 - newDotCount);
      });
    }, 300);

    return () => {
      clearInterval(interval);
    };
  }, []);

  return { waitingDots };
};
