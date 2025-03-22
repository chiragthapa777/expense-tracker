import { Portal } from "@gorhom/portal";
import BottomSheet, {
  BottomSheetBackdrop,
  BottomSheetView,
  SNAP_POINT_TYPE,
} from "@gorhom/bottom-sheet";

import React, { RefObject, useCallback, useState } from "react";
import { useColor } from "@/theme";

type Props = {
  ref: RefObject<BottomSheet>;
  snapPoints?: ("10%" | "25%" | "50%" | "70%" | "80%" | "90%" | "100%")[];
  onChange?: (index: number, position: number, type: SNAP_POINT_TYPE) => void;
  enablePanDownToClose?: boolean;
  children: React.ReactNode;
  onClose?: () => void;
};

export const CustomBottomSheet = React.forwardRef<
  React.ElementRef<typeof BottomSheet>,
  Props
>(
  (
    {
      snapPoints = ["25%", "50%", "90%"],
      onChange,
      enablePanDownToClose = true,
      children,
      onClose,
    },
    ref
  ) => {
    const color = useColor();
    const [isOpen, setIsOpen] = useState(false);
    const onSheetChange = useCallback((index: number) => {
      if (index > -1) {
        setIsOpen(true);
      } else {
        setIsOpen(false);
      }
    }, []);

    return (
      <Portal>
        <BottomSheet
          ref={ref}
          snapPoints={snapPoints}
          enableDynamicSizing={true}
          onChange={(
            index: number,
            position: number,
            type: SNAP_POINT_TYPE
          ) => {
            onSheetChange(index);
            if (onChange) {
              onChange(index, position, type);
            }
          }}
          enablePanDownToClose={enablePanDownToClose}
          index={-1}
          onClose={onClose}
          handleIndicatorStyle={{
            backgroundColor: color.border,
          }}
          handleStyle={{
            backgroundColor: color.background,
            borderTopEndRadius: 20,
            borderTopStartRadius: 20,
          }}
          backgroundStyle={{
            backgroundColor: color.background,
          }}
          backdropComponent={(props) => (
            <BottomSheetBackdrop
              {...props}
              disappearsOnIndex={-1}
              appearsOnIndex={0}
              pressBehavior="close"
            />
          )}
        >
          <BottomSheetView style={{ flex: 1 }}>{children}</BottomSheetView>
        </BottomSheet>
      </Portal>
    );
  }
);
