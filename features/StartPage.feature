@future
Feature: Start Page

  Scenario: Start Page is present
    When accessing "/"
    Then the response Status-Code MUST be 200

  Scenario: The start page contains an explanation
    When accessing "/"
    Then the response body MUST be:
     """
     TODO
     """