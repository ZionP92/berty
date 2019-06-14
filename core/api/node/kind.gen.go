


// file generated by protoc-gen-gotemplate

package node

import (
  "fmt"
  "encoding/json"
)

      // GetNodeStartedAttrs is a typesafe version of GetAttrs
      func (e *NodeEvent) GetNodeStartedAttrs() (*NodeStartedAttrs, error) {
        if e.Attributes == nil || len(e.Attributes) == 0 {
          return &NodeStartedAttrs{}, nil
        }
        var attrs NodeStartedAttrs
        return &attrs, json.Unmarshal(e.Attributes, &attrs)
      }

      // SetNodeStartedAttrs is a typesafe version of the generic SetAttrs method
      func (e *NodeEvent) SetNodeStartedAttrs(attrs *NodeStartedAttrs) error {
        raw, err := json.Marshal(attrs)
        if err != nil {
          return err
        }
        e.Attributes = raw
        return nil
      }

      // GetNodeStoppedAttrs is a typesafe version of GetAttrs
      func (e *NodeEvent) GetNodeStoppedAttrs() (*NodeStoppedAttrs, error) {
        if e.Attributes == nil || len(e.Attributes) == 0 {
          return &NodeStoppedAttrs{}, nil
        }
        var attrs NodeStoppedAttrs
        return &attrs, json.Unmarshal(e.Attributes, &attrs)
      }

      // SetNodeStoppedAttrs is a typesafe version of the generic SetAttrs method
      func (e *NodeEvent) SetNodeStoppedAttrs(attrs *NodeStoppedAttrs) error {
        raw, err := json.Marshal(attrs)
        if err != nil {
          return err
        }
        e.Attributes = raw
        return nil
      }

      // GetNodeIsAliveAttrs is a typesafe version of GetAttrs
      func (e *NodeEvent) GetNodeIsAliveAttrs() (*NodeIsAliveAttrs, error) {
        if e.Attributes == nil || len(e.Attributes) == 0 {
          return &NodeIsAliveAttrs{}, nil
        }
        var attrs NodeIsAliveAttrs
        return &attrs, json.Unmarshal(e.Attributes, &attrs)
      }

      // SetNodeIsAliveAttrs is a typesafe version of the generic SetAttrs method
      func (e *NodeEvent) SetNodeIsAliveAttrs(attrs *NodeIsAliveAttrs) error {
        raw, err := json.Marshal(attrs)
        if err != nil {
          return err
        }
        e.Attributes = raw
        return nil
      }

      // GetBackgroundCriticalAttrs is a typesafe version of GetAttrs
      func (e *NodeEvent) GetBackgroundCriticalAttrs() (*BackgroundCriticalAttrs, error) {
        if e.Attributes == nil || len(e.Attributes) == 0 {
          return &BackgroundCriticalAttrs{}, nil
        }
        var attrs BackgroundCriticalAttrs
        return &attrs, json.Unmarshal(e.Attributes, &attrs)
      }

      // SetBackgroundCriticalAttrs is a typesafe version of the generic SetAttrs method
      func (e *NodeEvent) SetBackgroundCriticalAttrs(attrs *BackgroundCriticalAttrs) error {
        raw, err := json.Marshal(attrs)
        if err != nil {
          return err
        }
        e.Attributes = raw
        return nil
      }

      // GetBackgroundErrorAttrs is a typesafe version of GetAttrs
      func (e *NodeEvent) GetBackgroundErrorAttrs() (*BackgroundErrorAttrs, error) {
        if e.Attributes == nil || len(e.Attributes) == 0 {
          return &BackgroundErrorAttrs{}, nil
        }
        var attrs BackgroundErrorAttrs
        return &attrs, json.Unmarshal(e.Attributes, &attrs)
      }

      // SetBackgroundErrorAttrs is a typesafe version of the generic SetAttrs method
      func (e *NodeEvent) SetBackgroundErrorAttrs(attrs *BackgroundErrorAttrs) error {
        raw, err := json.Marshal(attrs)
        if err != nil {
          return err
        }
        e.Attributes = raw
        return nil
      }

      // GetBackgroundWarnAttrs is a typesafe version of GetAttrs
      func (e *NodeEvent) GetBackgroundWarnAttrs() (*BackgroundWarnAttrs, error) {
        if e.Attributes == nil || len(e.Attributes) == 0 {
          return &BackgroundWarnAttrs{}, nil
        }
        var attrs BackgroundWarnAttrs
        return &attrs, json.Unmarshal(e.Attributes, &attrs)
      }

      // SetBackgroundWarnAttrs is a typesafe version of the generic SetAttrs method
      func (e *NodeEvent) SetBackgroundWarnAttrs(attrs *BackgroundWarnAttrs) error {
        raw, err := json.Marshal(attrs)
        if err != nil {
          return err
        }
        e.Attributes = raw
        return nil
      }

      // GetDebugAttrs is a typesafe version of GetAttrs
      func (e *NodeEvent) GetDebugAttrs() (*DebugAttrs, error) {
        if e.Attributes == nil || len(e.Attributes) == 0 {
          return &DebugAttrs{}, nil
        }
        var attrs DebugAttrs
        return &attrs, json.Unmarshal(e.Attributes, &attrs)
      }

      // SetDebugAttrs is a typesafe version of the generic SetAttrs method
      func (e *NodeEvent) SetDebugAttrs(attrs *DebugAttrs) error {
        raw, err := json.Marshal(attrs)
        if err != nil {
          return err
        }
        e.Attributes = raw
        return nil
      }

      // GetStatisticsAttrs is a typesafe version of GetAttrs
      func (e *NodeEvent) GetStatisticsAttrs() (*StatisticsAttrs, error) {
        if e.Attributes == nil || len(e.Attributes) == 0 {
          return &StatisticsAttrs{}, nil
        }
        var attrs StatisticsAttrs
        return &attrs, json.Unmarshal(e.Attributes, &attrs)
      }

      // SetStatisticsAttrs is a typesafe version of the generic SetAttrs method
      func (e *NodeEvent) SetStatisticsAttrs(attrs *StatisticsAttrs) error {
        raw, err := json.Marshal(attrs)
        if err != nil {
          return err
        }
        e.Attributes = raw
        return nil
      }

  // GetAttrs parses the embedded attributes
  func (e *NodeEvent) GetAttrs() (interface{}, error) {
    switch e.Kind {
          case Kind_NodeStarted:
            return e.GetNodeStartedAttrs()
          case Kind_NodeStopped:
            return e.GetNodeStoppedAttrs()
          case Kind_NodeIsAlive:
            return e.GetNodeIsAliveAttrs()
          case Kind_BackgroundCritical:
            return e.GetBackgroundCriticalAttrs()
          case Kind_BackgroundError:
            return e.GetBackgroundErrorAttrs()
          case Kind_BackgroundWarn:
            return e.GetBackgroundWarnAttrs()
          case Kind_Debug:
            return e.GetDebugAttrs()
          case Kind_Statistics:
            return e.GetStatisticsAttrs()
    }
    return nil, fmt.Errorf("not supported event kind: %q", e.Kind)
  }


















