import { useEffect, useState } from "react";

export const useWaitingDots = () => {
  const [waitingDots, setWaitingDots] = useState("   ");

  useEffect(() => {
    const interval = setInterval(() => {
      setWaitingDots((prevDots) => {
        const dotCount = prevDots.replace(/\s/g, "").length;

        if (dotCount === 3) {
          return "   "; 
        }

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
