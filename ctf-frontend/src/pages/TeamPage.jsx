
import Stepper, { Step } from "../components/Stepper/Stepper";
import LetterGlitch from "../components/background/LetterGlitch";

import TeamActionStep from "./team/TeamActionStep";

export default function TeamPage() {
  return (
    <div style={{ height: "100vh", position: "relative" }}>
      <LetterGlitch />

      <div
        style={{
          position: "absolute",
          inset: 0,
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
        }}
      >
        
          <Stepper
            nextButtonText="Continue"
            onFinalStepCompleted={() => {
              console.log("Team setup completed");
              // later → navigate to dashboard
            }}
          >
            <Step>
              <TeamActionStep />
            </Step>
          </Stepper>
    
      </div>
    </div>
  );
}
