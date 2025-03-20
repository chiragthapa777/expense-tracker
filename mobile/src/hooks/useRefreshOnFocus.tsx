import { useFocusEffect } from "@react-navigation/native";
import React, { useRef } from "react";

export function useRefreshOnFocus<T>(refetch: () => Promise<T>) {
  const firstTimeRef = useRef(true);
  useFocusEffect(
    React.useCallback(() => {
      if (firstTimeRef.current) {
        firstTimeRef.current = false;
        return;
      }
      refetch();
    }, [refetch])
  );
}

// Example
// const Expenses = () => {
//   const isFocused = useIsFocused();
//   // Fetch expenses (Read)
//   const {
//     data: expenses,
//     isLoading,
//     error,
//     refetch,
//   } = useQuery<Expense[], Error>({
//     queryKey: ["expenses"],
//     queryFn: fetchExpenses,
//     subscribed: isFocused, // Only subscribe when screen is focused
//   });
//   // Refresh on screen focus
//   useRefreshOnFocus(refetch);
// };
